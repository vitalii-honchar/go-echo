package middleware

import (
	"go-echo/internal/lib"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(asMiddleware(NewLoggingMiddleware)),
)

func asMiddleware(f any) any {
	return lib.AsGroup(f, new(Middleware), "middleware")
}
