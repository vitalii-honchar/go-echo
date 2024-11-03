package internal

import (
	"go-echo/internal/infra/config"
	"go-echo/internal/infra/web"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewApp() *fx.App {
	return fx.New(
		web.Module,
		fx.Provide(
			zap.NewDevelopment,
			config.NewConfig,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
}
