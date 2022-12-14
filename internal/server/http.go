package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/service"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"gitlab.com/mcsolutions/find-psy/back/common/config/conf_servers"
	pb "gitlab.com/mcsolutions/find-psy/proto/gen/go/api/article"
	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

// NewHTTPServer new HTTP server.
func NewHTTPServer(confServer *common.Server, article *service.ArticleService, logger log.Logger) (*http.Server, error) {
	if confServer == nil {
		return nil, cerrors.GetBadConfigError("server")
	}
	opts, err := conf_servers.GetHttpOpts(confServer.Http, logger)
	if err != nil {
		return nil, err
	}
	srv := http.NewServer(opts...)
	srv.Handle("/metrics", promhttp.Handler())
	pb.RegisterArticleHTTPServer(srv, article)
	return srv, nil
}
