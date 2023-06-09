package main

import (
	"context"
	"os"
	"time"

	"github.com/go-jimu/components/config/loader"
	"github.com/go-jimu/components/mediator"
	"github.com/go-jimu/components/sloghelper"
	"github.com/jacexh/linear-app-intergation/internal/bootstrap"
	"github.com/jacexh/linear-app-intergation/internal/bootstrap/httpsrv"
	"github.com/jacexh/linear-app-intergation/internal/bootstrap/mysql"
	"github.com/jacexh/linear-app-intergation/internal/business/user"
	"go.uber.org/fx"
	"golang.org/x/exp/slog"
)

type Option struct {
	fx.Out
	Logger     sloghelper.Options `json:"logger" toml:"logger" yaml:"logger"`
	MySQL      mysql.Option       `json:"mysql" toml:"mysql" yaml:"mysql"`
	HTTPServer httpsrv.Option     `json:"http-server" toml:"http-server" yaml:"http-server"`
	Eventbus   mediator.Options   `json:"eventbus" toml:"eventbus" yaml:"eventbus"`
}

func parseOption() (Option, error) {
	opt := new(Option)
	err := loader.Load(opt)
	return *opt, err
}

func main() {
	app := fx.New(
		fx.Provide(parseOption),
		fx.Provide(sloghelper.NewLog),
		fx.Provide(mediator.NewInMemMediator),
		fx.Invoke(mediator.SetDefault),
		bootstrap.Module,
		user.Module,
		fx.NopLogger,
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Start(startCtx); err != nil {
		slog.ErrorCtx(startCtx, "启动应用失败", sloghelper.Error(err))
		os.Exit(1)
	}

	<-app.Done()
	slog.Warn("caught quit signal")

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		slog.Error("failed to stop application", sloghelper.Error(err))
		os.Exit(1)
	}

	slog.Info("bye")
	os.Exit(0)
}
