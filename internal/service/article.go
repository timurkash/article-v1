package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/biz"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"gitlab.com/mcsolutions/find-psy/back/common/paging"
	pb "gitlab.com/mcsolutions/find-psy/proto/gen/go/api/article"
	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

type ArticleService struct {
	uc  *biz.ArticleUsecase
	log *log.Helper

	pb.UnimplementedArticleServer
}

func NewArticleService(uc *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *ArticleService) Create(ctx context.Context, req *pb.CreateArticleRequest) (*pb.ArticleReply, error) {
	if req.Article == nil {
		return nil, cerrors.GetRequiredError("article")
	}
	if req.Article.Title == "" {
		return nil, cerrors.GetRequiredError("article.title")
	}
	if req.Article.Html == "" {
		return nil, cerrors.GetRequiredError("article.text")
	}
	if len(req.Article.Tags) == 0 {
		return nil, cerrors.GetRequiredError("article.tags")
	}
	if err := cerrors.CheckLang(req.Article.Lang); err != nil {
		return nil, err
	}
	article, err := s.uc.Create(ctx, req.Article)
	if err != nil {
		return nil, err
	}
	return &pb.ArticleReply{
		Article: article,
	}, nil
}

func (s *ArticleService) Update(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.ArticleReply, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	if req.Article == nil {
		return nil, cerrors.GetRequiredError("article")
	}
	if req.Article.Title == "" {
		return nil, cerrors.GetRequiredError("article.title")
	}
	if req.Article.Html == "" {
		return nil, cerrors.GetRequiredError("article.text")
	}
	if len(req.Article.Tags) == 0 {
		return nil, cerrors.GetRequiredError("article.tags")
	}
	article, err := s.uc.Update(ctx, id, req.Article)
	if err != nil {
		return nil, err
	}
	return &pb.ArticleReply{
		Article: article,
	}, nil
}

func (s *ArticleService) Delete(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	if err := s.uc.Delete(ctx, id); err != nil {
		return nil, err
	}
	return &common.Empty{}, nil
}

func (s *ArticleService) Get(ctx context.Context, req *common.IdRequest) (*pb.ArticleReply, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	article, err := s.uc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.ArticleReply{
		Article: article,
	}, nil
}

func (s *ArticleService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListReply, error) {
	if req.Filter == nil {
		return nil, cerrors.GetRequiredError("filter")
	}
	articles, thePaging, err := s.uc.List(ctx, req.Filter, paging.GetOrderOffsetLimit(req.Ool))
	if err != nil {
		return nil, err
	}
	return &pb.ListReply{
		Articles: articles,
		Paging:   thePaging.GetReply(),
	}, nil
}

func (s *ArticleService) ListMy(ctx context.Context, req *pb.ListMyRequest) (*pb.ListReply, error) {
	articles, thePaging, err := s.uc.ListMy(ctx, req.Published, paging.GetOrderOffsetLimit(req.Ool))
	if err != nil {
		return nil, err
	}
	return &pb.ListReply{
		Articles: articles,
		Paging:   thePaging.GetReply(),
	}, nil
}

func (s *ArticleService) Publish(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	if err := s.uc.Publish(ctx, id); err != nil {
		return nil, err
	}
	return &common.Empty{}, nil
}

func (s *ArticleService) Revoke(ctx context.Context, req *common.IdRequest) (*common.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	if err := s.uc.Revoke(ctx, id); err != nil {
		return nil, err
	}
	return &common.Empty{}, nil
}

func (s *ArticleService) SetProcessed(ctx context.Context, req *pb.SetProcessedRequest) (*common.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	if req.Html == "" {
		return nil, cerrors.GetRequiredError("html")
	}
	if len(req.Tags) == 0 {
		return nil, cerrors.GetRequiredError("tags")
	}
	if err := s.uc.SetProcessed(ctx, id, req.Html, req.Tags); err != nil {
		return nil, err
	}
	return &common.Empty{}, nil
}

func (s *ArticleService) ListToProcess(ctx context.Context, req *pb.ListToProcessRequest) (*pb.ListReply, error) {
	if req.Filter == nil {
		return nil, cerrors.GetRequiredError("filter")
	}
	articles, thePaging, err := s.uc.ListToProcess(ctx,
		req.Filter,
		paging.GetOrderOffsetLimit(req.Ool),
	)
	if err != nil {
		return nil, err
	}
	return &pb.ListReply{
		Articles: articles,
		Paging:   thePaging.GetReply(),
	}, nil
}
