package user

import (
	"github.com/jacexh/linear-app-intergation/internal/bootstrap/httpsrv"
	"github.com/jacexh/linear-app-intergation/internal/business/user/application"
	"github.com/jacexh/linear-app-intergation/internal/business/user/application/transport"
	"github.com/jacexh/linear-app-intergation/internal/business/user/infrastructure"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"domain.user",
	fx.Provide(infrastructure.NewQueryRepository),
	fx.Provide(transport.NewController),
	fx.Provide(application.NewApplication),
	fx.Provide(infrastructure.NewRepository),
	fx.Invoke(func(srv httpsrv.HTTPServer, controller httpsrv.Controller) {
		srv.With(controller)
	}),
)
