package bootstrap

import (
	"github.com/jacexh/linear-app-intergation/internal/bootstrap/httpsrv"
	"github.com/jacexh/linear-app-intergation/internal/bootstrap/mysql"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"bootstrap",
	fx.Provide(httpsrv.NewHTTPServer),
	fx.Provide(mysql.NewMySQLDriver),
)
