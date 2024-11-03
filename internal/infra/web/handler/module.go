package handler

import (
	"go-echo/internal/infra/web/handler/user"
	"go-echo/internal/lib"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(asHandler(user.NewGetUserHandler)),
)

func asHandler(f any) any {
	return lib.AsGroup(f, new(Handler), "handler")
}
