package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/conf"
	"gitlab.com/mcsolutions/find-psy/back/common/ccontext"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"gitlab.com/mcsolutions/find-psy/back/common/jwt"
	"gitlab.com/mcsolutions/find-psy/back/common/others"
	"gitlab.com/mcsolutions/find-psy/back/common/paging"
	pb "gitlab.com/mcsolutions/find-psy/proto/gen/go/api/article"
	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

type (
	ArticleRepo interface {
		Create(context.Context, uuid.UUID, *pb.ArticleInfo) error
		Update(context.Context, uuid.UUID, *pb.ArticleInfo) error
		Delete(context.Context, uuid.UUID) error
		Get(context.Context, uuid.UUID) (*pb.ArticleInfo, error)
		SetPublished(context.Context, uuid.UUID, bool) error
		List(context.Context, *pb.Filter, *paging.OrderOffsetLimit) ([]*pb.ArticleInfo, *paging.Paging, error)
		ListMy(context.Context, string, *common.Bool, *paging.OrderOffsetLimit) ([]*pb.ArticleInfo, *paging.Paging, error)
		ListToProcess(context.Context, *pb.FilterToProcess, *paging.OrderOffsetLimit) ([]*pb.ArticleInfo, *paging.Paging, error)
		SetProcessed(context.Context, uuid.UUID, string, []string) error
	}
	ICache interface {
		SetEx(context.Context, string, time.Duration) error
	}
)

type ArticleUsecase struct {
	publishDuration time.Duration
	shortenLimit    int32
	repo            ArticleRepo
	cache           ICache
	log             *log.Helper
}

func NewArticleUsecase(confBusiness *conf.Business, repo ArticleRepo, cache ICache, logger log.Logger) (*ArticleUsecase, error) {
	if confBusiness == nil {
		return nil, cerrors.GetBadConfigError("business")
	}
	if confBusiness.PublishDuration == nil {
		return nil, cerrors.GetBadConfigError("business.publish_duration")
	}
	var shortenLimit int32 = 1000
	if confBusiness.ShortenLimit > 0 {
		shortenLimit = confBusiness.ShortenLimit
	}
	return &ArticleUsecase{
		publishDuration: confBusiness.PublishDuration.AsDuration(),
		shortenLimit:    shortenLimit,
		repo:            repo,
		cache:           cache,
		log:             log.NewHelper(logger),
	}, nil
}

func (uc *ArticleUsecase) Create(ctx context.Context, in *pb.InArticleInfo) (*pb.ArticleInfo, error) {
	if ccontext.IsNotAuthorized(ctx) {
		return nil, cerrors.NotAuthorized
	}
	claims := jwt.GetClaims(ctx)
	if claims.UserType == "" {
		return nil, cerrors.IsNotPsyAndAdmin
	}
	if claims.UserType == jwt.Admin {
		if !others.Contains(claims.RealmAccess.Roles, jwt.ArticleAdmin) {
			return nil, cerrors.GetIsNotAdminError(jwt.ArticleAdmin)
		}
	}
	id := uuid.New()
	articleInfo := &pb.ArticleInfo{
		Id:         id.String(),
		Email:      claims.Email,
		UserName:   claims.UserName,
		Title:      in.Title,
		Latinized:  others.GetLatinized(in.Title),
		Html:       in.Html,
		ShortText:  GetTextAndShort(in.Html, int(uc.shortenLimit)),
		Lang:       in.Lang,
		Tags:       in.Tags,
		Cdate:      timestamppb.Now(),
		Source:     in.Source,
		B17Account: in.B17Account,
	}
	if claims.UserType == jwt.Admin {
		articleInfo.IsPublished = true
		articleInfo.PubDate = timestamppb.Now()
	}
	if err := uc.repo.Create(ctx, id, articleInfo); err != nil {
		return nil, err
	}
	return articleInfo, nil
}

func (uc *ArticleUsecase) Update(ctx context.Context, id uuid.UUID, in *pb.InArticleInfo) (*pb.ArticleInfo, error) {
	if ccontext.IsNotAuthorized(ctx) {
		return nil, cerrors.NotAuthorized
	}
	claims := jwt.GetClaims(ctx)
	if claims.UserType != jwt.Psy {
		return nil, cerrors.IsNotPsy
	}
	article, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if article.IsPublished {
		if err := uc.cache.SetEx(ctx, claims.Email, uc.publishDuration); err != nil {
			return nil, err
		}
	}
	if article.Email != claims.Email {
		return nil, cerrors.NotPermitted
	}
	article.Title = in.Title
	article.Latinized = others.GetLatinized(article.Title)
	article.Html = in.Html
	article.ShortText = GetTextAndShort(in.Html, int(uc.shortenLimit))
	article.Lang = in.Lang
	article.Udate = timestamppb.Now()
	if err := uc.repo.Update(ctx, id, article); err != nil {
		return nil, err
	}
	return uc.repo.Get(ctx, id)
}

func (uc *ArticleUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	if ccontext.IsNotAuthorized(ctx) {
		return cerrors.NotAuthorized
	}
	claims := jwt.GetClaims(ctx)
	if claims.UserType == "" {
		return cerrors.IsNotPsyAndAdmin
	}
	switch claims.UserType {
	case jwt.Psy:
		article, err := uc.Get(ctx, id)
		if err != nil {
			return err
		}
		if article.Email != claims.Email {
			return cerrors.NotPermitted
		}
		if err := uc.Delete(ctx, id); err != nil {
			return err
		}
	case jwt.Admin:
		if !others.Contains(claims.RealmAccess.Roles, jwt.PsyAdmin) {
			return cerrors.GetIsNotAdminError(jwt.PsyAdmin)
		}
		if err := uc.Delete(ctx, id); err != nil {
			return err
		}
	}
	return uc.repo.Delete(ctx, id)
}

func (uc *ArticleUsecase) Get(ctx context.Context, id uuid.UUID) (*pb.ArticleInfo, error) {
	short := ccontext.IsNotAuthorized(ctx)
	article, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if short {
		article.Html = shortenHtml(article.Html, uc.shortenLimit)
	}
	return article, nil
}

func (uc *ArticleUsecase) List(ctx context.Context,
	filter *pb.Filter,
	ool *paging.OrderOffsetLimit,
) ([]*pb.ArticleInfo, *paging.Paging, error) {
	short := ccontext.IsNotAuthorized(ctx)
	articles, thePaging, err := uc.repo.List(ctx, filter, ool)
	if err != nil {
		return nil, nil, err
	}
	if short {
		for i := range articles {
			articles[i].Html = shortenHtml(articles[i].Html, uc.shortenLimit)
		}
	}
	return articles, thePaging, nil
}

func (uc *ArticleUsecase) ListMy(ctx context.Context,
	isPublished *common.Bool,
	ool *paging.OrderOffsetLimit,
) ([]*pb.ArticleInfo, *paging.Paging, error) {
	if ccontext.IsNotAuthorized(ctx) {
		return nil, nil, cerrors.NotAuthorized
	}
	claims := jwt.GetClaims(ctx)
	if claims.UserType != jwt.Psy {
		return nil, nil, cerrors.IsNotPsy
	}
	return uc.repo.ListMy(ctx, claims.Email, isPublished, ool)
}

func (uc *ArticleUsecase) Publish(ctx context.Context, id uuid.UUID) error {
	if ccontext.IsNotAuthorized(ctx) {
		return cerrors.NotAuthorized
	}
	claims := jwt.GetClaims(ctx)
	if claims.UserType != jwt.Psy {
		return cerrors.IsNotPsy
	}
	if err := uc.cache.SetEx(ctx, claims.Email, uc.publishDuration); err != nil {
		return err
	}
	article, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}
	if article.Email != claims.Email {
		return cerrors.NotPermitted
	}
	if article.IsPublished {
		return cerrors.AlreadyPublished
	}
	article.IsPublished = true
	article.PubDate = timestamppb.Now()
	if err := uc.repo.SetPublished(ctx, id, true); err != nil {
		return err
	}
	return nil
}

func (uc *ArticleUsecase) Revoke(ctx context.Context, id uuid.UUID) error {
	if ccontext.IsNotAuthorized(ctx) {
		return cerrors.NotAuthorized
	}
	claims := jwt.GetClaims(ctx)
	if claims.UserType != jwt.Psy {
		return cerrors.IsNotPsy
	}
	article, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}
	if article.Email != claims.Email {
		return cerrors.NotPermitted
	}
	if !article.IsPublished {
		return cerrors.NotPublished
	}
	if err := uc.repo.SetPublished(ctx, id, false); err != nil {
		return err
	}
	return nil
}

func (uc *ArticleUsecase) SetProcessed(ctx context.Context, id uuid.UUID, html string, tags []string) error {
	if !jwt.IsAdmin(ctx, jwt.ArticleAdmin) {
		return cerrors.IsNotAdmin
	}
	if err := uc.repo.SetProcessed(ctx, id, html, tags); err != nil {
		return err
	}
	return nil
}

func (uc *ArticleUsecase) ListToProcess(ctx context.Context,
	filter *pb.FilterToProcess,
	ool *paging.OrderOffsetLimit,
) ([]*pb.ArticleInfo, *paging.Paging, error) {
	if !jwt.IsAdmin(ctx, jwt.ArticleAdmin) {
		return nil, nil, cerrors.IsNotAdmin
	}
	articles, thePaging, err := uc.repo.ListToProcess(ctx, filter, ool)
	if err != nil {
		return nil, nil, err
	}
	return articles, thePaging, nil
}
