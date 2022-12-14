package conf_servers

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/sentry"
	"time"

	"gitlab.com/mcsolutions/find-psy/back/common/cerrors"
	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

func getMiddlewares(logger log.Logger, additionalMiddlewares ...middleware.Middleware) []middleware.Middleware {
	return append([]middleware.Middleware{
		recovery.Recovery(),
		sentry.Server(),
		tracing.Server(),
		logging.Server(logger),
	}, additionalMiddlewares...)
}

func GetGrpcOpts(grpcServerParams *common.ServerParams, logger log.Logger, middlewares ...middleware.Middleware) ([]grpc.ServerOption, error) {
	if grpcServerParams.Addr == "" {
		return nil, cerrors.GetBadConfigError("grpc.addr")
	}
	timeout := time.Second
	if grpcServerParams.Timeout != nil {
		timeout = grpcServerParams.Timeout.AsDuration()
	}
	opts := make([]grpc.ServerOption, 0, 4)
	opts = append(opts,
		grpc.Address(grpcServerParams.Addr),
		grpc.Timeout(timeout),
		grpc.Middleware(getMiddlewares(logger, middlewares...)...),
	)
	if grpcServerParams.Network != "" {
		opts = append(opts, grpc.Network(grpcServerParams.Network))
	}
	return opts, nil
}

func GetHttpOpts(httpServerParams *common.ServerParams, logger log.Logger, middlewares ...middleware.Middleware) ([]http.ServerOption, error) {
	if httpServerParams.Addr == "" {
		return nil, cerrors.GetBadConfigError("http.addr")
	}
	timeout := time.Second
	if httpServerParams.Timeout != nil {
		timeout = httpServerParams.Timeout.AsDuration()
	}
	opts := make([]http.ServerOption, 0, 4)
	opts = append(opts,
		http.Address(httpServerParams.Addr),
		http.Timeout(timeout),
		http.Middleware(getMiddlewares(logger, middlewares...)...),
	)
	if httpServerParams.Network != "" {
		opts = append(opts, http.Network(httpServerParams.Network))
	}
	return opts, nil
}
