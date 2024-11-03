package web

import (
	"go-echo/internal/infra/web/handler"
	"go-echo/internal/infra/web/middleware"

	"go.uber.org/fx"
)

var Module = fx.Options(
	handler.Module,
	middleware.Module,
	fx.Provide(newRenderer),
	fx.Provide(
		fx.Annotate(
			newServer,
			fx.ParamTags(`group:"middleware"`, `group:"handler"`),
		),
	),
	fx.Invoke(startServer),
)
