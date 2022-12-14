package service

import (
	"context"
	"os"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/biz"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/test"
)

func getTestService() (*ArticleService, context.Context, error) {
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	test.IsUnitTest = true
	dataData, _, err := data.NewData(nil, logger)
	if err != nil {
		return nil, nil, err
	}
	articleRepo := data.NewArticleRepo(dataData, logger)
	articleUsecase := biz.NewArticleUsecase(nil, articleRepo, logger)
	return NewArticleService(articleUsecase, logger), context.TODO(), nil
}

func Test_SayHello(t *testing.T) {
	service, ctx, err := getTestService()
	if err != nil {
		t.Fatal(err)
	}
	service = service
	ctx = ctx
	//{
	//	in := &pb.HelloRequest{
	//		Name: test.Test,
	//	}
	//	reply, err := service.SayHello(ctx, in)
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	if reply.Message != "Hello "+test.Test {
	//		t.Fatal("unexpected message")
	//	}
	//}
	//{
	//	in := &pb.HelloRequest{
	//		Name: test.Error,
	//	}
	//	_, err := service.SayHello(ctx, in)
	//	if err == nil {
	//		t.Fatal("expected error")
	//	}
	//}
}
