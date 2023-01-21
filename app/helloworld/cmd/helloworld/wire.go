//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/EndlessIdea/kratos-app/app/helloworld/internal/biz"
	"github.com/EndlessIdea/kratos-app/app/helloworld/internal/conf"
	"github.com/EndlessIdea/kratos-app/app/helloworld/internal/data"
	"github.com/EndlessIdea/kratos-app/app/helloworld/internal/server"
	"github.com/EndlessIdea/kratos-app/app/helloworld/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
