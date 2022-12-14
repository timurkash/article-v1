package data

import (
	"context"
	"fmt"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent/predicate"
	"gitlab.com/mcsolutions/find-psy/back/common/consts"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/biz"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data/ent/article"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/test"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"gitlab.com/mcsolutions/find-psy/back/common/others"
	"gitlab.com/mcsolutions/find-psy/back/common/paging"
	pb "gitlab.com/mcsolutions/find-psy/proto/gen/go/api/article"
	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) Create(ctx context.Context, id uuid.UUID, info *pb.ArticleInfo) error {
	articleCreate := r.data.relational.Article.Create().
		SetID(id).
		SetEmail(info.Email).
		SetUserName(info.UserName).
		SetTitle(info.Title).
		SetLatinized(info.Latinized).
		SetHTML(info.Html).
		SetShortText(info.ShortText).
		SetLang(info.Lang).
		SetTags(others.CommasAround(info.Tags)).
		SetIsPublished(info.IsPublished).
		SetCdate(time.Now())
	if info.B17Account != nil {
		articleCreate.SetB17Account(info.B17Account.Value)
	}
	if info.Source != nil {
		articleCreate.SetSource(info.Source.Value)
	}
	if info.PubDate != nil {
		articleCreate.SetPubDate(info.PubDate.AsTime())
	}
	if _, err := articleCreate.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (r *articleRepo) Update(ctx context.Context, id uuid.UUID, info *pb.ArticleInfo) error {
	if test.IsUnitTest {
		return nil
	}
	r.log.Info("article updating")
	updateArticleRecord := r.data.relational.Article.Update().Where(article.ID(id)).
		SetTitle(info.Title).
		SetLatinized(info.Latinized).
		SetHTML(info.Html).
		SetShortText(info.ShortText).
		SetLang(info.Lang).
		SetTags(others.CommasAround(info.Tags)).
		SetIsPublished(info.IsPublished).
		SetUdate(time.Now())
	if info.PubDate != nil {
		updateArticleRecord.SetPubDate(info.PubDate.AsTime())
	} else {
		updateArticleRecord.ClearPubDate()
	}
	if info.B17Account != nil {
		updateArticleRecord.SetB17Account(info.B17Account.Value)
	} else {
		updateArticleRecord.ClearB17Account()
	}
	if info.Source != nil {
		updateArticleRecord.SetSource(info.Source.Value)
	} else {
		updateArticleRecord.ClearSource()
	}
	count, err := updateArticleRecord.Save(ctx)
	if err != nil {
		return err
	}
	if count != 1 {
		return cerrors.NotFound
	}
	return nil
}

func (r *articleRepo) SetPublished(ctx context.Context, id uuid.UUID, published bool) error {
	updateArticleRecord := r.data.relational.Article.Update().Where(article.ID(id)).
		SetIsPublished(published)
	if published {
		updateArticleRecord.SetPubDate(time.Now())
	} else {
		updateArticleRecord.ClearPubDate()
	}
	count, err := updateArticleRecord.Save(ctx)
	if err != nil {
		return err
	}
	if count == 0 {
		return cerrors.NotFound
	}
	return nil
}

func (r *articleRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if test.IsUnitTest {
		return nil
	}
	r.log.Info("article deleting")
	count, err := r.data.relational.Article.Delete().Where(article.ID(id)).Exec(ctx)
	if err != nil {
		return err
	}
	if count == 0 {
		return cerrors.NotFound
	}
	return nil
}

func (r *articleRepo) Get(ctx context.Context, id uuid.UUID) (*pb.ArticleInfo, error) {
	r.log.Info("article getting")
	articleRecord, err := r.data.relational.Article.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.ArticleInfo{
		Id:          articleRecord.ID.String(),
		Email:       articleRecord.Email,
		UserName:    articleRecord.UserName,
		Title:       articleRecord.Title,
		Latinized:   articleRecord.Latinized,
		Html:        articleRecord.HTML,
		ShortText:   articleRecord.ShortText,
		Tags:        others.PStringSplitSliceTrimCommas(articleRecord.Tags),
		Lang:        articleRecord.Lang,
		B17Account:  others.StringValue(articleRecord.B17Account),
		Source:      others.StringValue(articleRecord.Source),
		IsPublished: articleRecord.IsPublished,
		PubDate:     others.PTimestamp(articleRecord.PubDate),
		Cdate:       others.PTimestamp(&articleRecord.Cdate),
		Udate:       others.PTimestamp(articleRecord.Udate),
		IsProcessed: articleRecord.IsProcessed,
	}, nil
}

func (r *articleRepo) List(
	ctx context.Context,
	filter *pb.Filter,
	ool *paging.OrderOffsetLimit,
) ([]*pb.ArticleInfo, *paging.Paging, error) {
	articleQuery := r.data.relational.Article.Query().Where(article.IsPublished(true))
	if filter.Title != nil {
		articleQuery = articleQuery.Where(article.TitleContains(filter.Title.Value))
	}
	if len(filter.Tags) > 0 {
		var predicates []predicate.Article
		for _, tag := range filter.Tags {
			predicates = append(predicates, article.TagsContains(fmt.Sprintf(",%s,", tag)))
		}
		articleQuery.Where(article.Or(predicates...))
	}
	if filter.B17Account != nil {
		articleQuery = articleQuery.Where(article.B17Account(filter.B17Account.Value))
	}
	if filter.Lang != nil {
		articleQuery = articleQuery.Where(article.Lang(filter.B17Account.Value))
	}
	var (
		offset uint32
		limit  uint32
		total  int
		err    error
	)
	total, err = articleQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	if ool != nil {
		offset = ool.Offset
		limit = ool.Limit
		articleQuery.Offset(int(offset)).Limit(int(limit))
	}
	articleQuery.Order(ent.Desc(consts.PubDate), ent.Desc(consts.CDate))
	items, err := articleQuery.All(ctx)
	if err != nil {
		return nil, nil, err
	}
	articles := make([]*pb.ArticleInfo, 0)
	for _, item := range items {
		articles = append(articles, &pb.ArticleInfo{
			Id:          item.ID.String(),
			Email:       item.Email,
			UserName:    item.UserName,
			Title:       item.Title,
			Latinized:   item.Latinized,
			Html:        item.HTML,
			ShortText:   item.ShortText,
			Lang:        item.Lang,
			Tags:        others.PStringSplitSliceTrimCommas(item.Tags),
			B17Account:  others.StringValue(item.B17Account),
			Source:      others.StringValue(item.Source),
			IsPublished: item.IsPublished,
			PubDate:     others.PTimestamp(item.PubDate),
			Cdate:       others.PTimestamp(&item.Cdate),
			Udate:       others.PTimestamp(item.Udate),
			IsProcessed: item.IsProcessed,
		})
	}
	return articles, &paging.Paging{
		Order:   "pubdate desc, cdate desc",
		Offset:  ool.Offset,
		Limit:   ool.Limit,
		Total:   total,
		HasNext: len(articles)+int(ool.Offset) < total,
	}, nil
}

func (r *articleRepo) ListMy(ctx context.Context,
	email string,
	published *common.Bool,
	ool *paging.OrderOffsetLimit,
) ([]*pb.ArticleInfo, *paging.Paging, error) {
	articleQuery := r.data.relational.Article.Query().Where(article.Email(email))
	if published != nil {
		articleQuery.Where(article.IsPublished(published.Value))
	}
	var (
		offset uint32
		limit  uint32
		total  int
		err    error
	)
	total, err = articleQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	if ool != nil {
		offset = ool.Offset
		limit = ool.Limit
		articleQuery.Offset(int(offset)).Limit(int(limit))
	}
	articleQuery.Order(ent.Desc(consts.PubDate), ent.Desc(consts.CDate))
	items, err := articleQuery.All(ctx)
	if err != nil {
		return nil, nil, err
	}
	articles := make([]*pb.ArticleInfo, 0)
	for _, item := range items {
		articles = append(articles, &pb.ArticleInfo{
			Id:          item.ID.String(),
			Email:       item.Email,
			UserName:    item.UserName,
			Title:       item.Title,
			Latinized:   item.Latinized,
			Html:        item.HTML,
			ShortText:   item.ShortText,
			Lang:        item.Lang,
			Tags:        others.PStringSplitSliceTrimCommas(item.Tags),
			B17Account:  others.StringValue(item.B17Account),
			Source:      others.StringValue(item.Source),
			IsPublished: item.IsPublished,
			PubDate:     others.PTimestamp(item.PubDate),
			Cdate:       others.PTimestamp(&item.Cdate),
			Udate:       others.PTimestamp(item.Udate),
			IsProcessed: item.IsProcessed,
		})
	}
	return articles, &paging.Paging{
		Order:   "pubdate desc, cdate desc",
		Offset:  ool.Offset,
		Limit:   ool.Limit,
		Total:   total,
		HasNext: len(articles)+int(ool.Offset) < total,
	}, nil
}

func (r *articleRepo) ListToProcess(ctx context.Context,
	filter *pb.FilterToProcess,
	ool *paging.OrderOffsetLimit,
) ([]*pb.ArticleInfo, *paging.Paging, error) {
	articleQuery := r.data.relational.Article.Query().
		Where(article.IsProcessed(filter.IsProcessed))
	if filter.Title != nil {
		articleQuery.Where(article.TitleContains(filter.Title.Value))
	}
	if filter.Date != nil {
		d, err := time.Parse("02.01.06", filter.Date.Value)
		if err != nil {
			return nil, nil, err
		}
		articleQuery.Where(article.PubDateGTE(d), article.PubDateLT(d.Add(24*time.Hour)))
	}
	if filter.B17Account != nil {
		articleQuery.Where(article.B17Account(filter.B17Account.Value))
	}
	if filter.Tag != nil {
		articleQuery.Where(article.TagsContains(fmt.Sprintf(",%s,", filter.Tag)))
	}
	var (
		offset uint32
		limit  uint32
		total  int
		err    error
	)
	total, err = articleQuery.Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	if ool != nil {
		offset = ool.Offset
		limit = ool.Limit
		articleQuery.Offset(int(offset)).Limit(int(limit))
	}
	articleQuery.Order(ent.Desc(consts.PubDate))
	items, err := articleQuery.All(ctx)
	if err != nil {
		return nil, nil, err
	}
	articles := make([]*pb.ArticleInfo, 0)
	for _, item := range items {
		articles = append(articles, &pb.ArticleInfo{
			Id:          item.ID.String(),
			Email:       item.Email,
			UserName:    item.UserName,
			Title:       item.Title,
			Latinized:   item.Latinized,
			Html:        item.HTML,
			ShortText:   item.ShortText,
			Lang:        item.Lang,
			Tags:        others.PStringSplitSliceTrimCommas(item.Tags),
			B17Account:  others.StringValue(item.B17Account),
			Source:      others.StringValue(item.Source),
			IsPublished: item.IsPublished,
			PubDate:     others.PTimestamp(item.PubDate),
			Cdate:       others.PTimestamp(&item.Cdate),
			Udate:       others.PTimestamp(item.Udate),
			IsProcessed: item.IsProcessed,
		})
	}
	return articles, &paging.Paging{
		Order:   "pubdate desc",
		Offset:  ool.Offset,
		Limit:   ool.Limit,
		Total:   total,
		HasNext: len(articles)+int(ool.Offset) < total,
	}, nil
}

func (r *articleRepo) SetProcessed(ctx context.Context, id uuid.UUID, html string, tags []string) error {
	count, err := r.data.relational.Article.Update().Where(article.ID(id)).
		SetIsPublished(true).
		SetIsProcessed(true).
		SetHTML(html).
		SetTags(others.CommasAround(tags)).
		Save(ctx)
	if err != nil {
		return err
	}
	if count == 0 {
		return cerrors.NotFound
	}
	return nil
}
