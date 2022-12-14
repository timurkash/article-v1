//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/biz"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/conf"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/outside/data"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/server"
	"gitlab.com/mcsolutions/find-psy/back/article-v1/internal/service"
	"gitlab.com/mcsolutions/find-psy/proto/gen/go/common"
)

// initApp init kratos application.
func initApp(
	*common.Server,
	*common.Jwks,
	*conf.Business,
	*conf.Data,
	log.Logger,
) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			service.ProviderSet,
			biz.ProviderSet,
			data.ProviderSet,
			newApp,
		),
	)
}
