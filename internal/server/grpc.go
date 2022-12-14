package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/service"
	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"gitlab.com/mcsolutions/find-psy/back/common/config/conf_jwks"
	"gitlab.com/mcsolutions/find-psy/back/common/config/conf_servers"
	"gitlab.com/mcsolutions/find-psy/back/common/jwt"
	pb "gitlab.com/mcsolutions/find-psy/proto/gen/go/api/article"
	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(confServer *common.Server, confJwks *common.Jwks, article *service.ArticleService, logger log.Logger) (*grpc.Server, error) {
	if confServer == nil {
		return nil, cerrors.GetBadConfigError("server")
	}
	jwks, err := conf_jwks.GetJWKS(confJwks)
	if err != nil {
		return nil, err
	}
	opts, err := conf_servers.GetGrpcOpts(confServer.Grpc, logger, jwt.GetAuthMiddleware(jwks))
	if err != nil {
		return nil, err
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterArticleServer(srv, article)
	return srv, nil
}
