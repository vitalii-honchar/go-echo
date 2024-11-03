package web

import (
	"context"
	"fmt"
	"go-echo/internal/infra/config"
	"go-echo/internal/infra/web/handler"
	"go-echo/internal/infra/web/middleware"
	"net/http"

	"github.com/h3poteto/pongo2echo"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"

	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func newServer(
	middlewares []middleware.Middleware,
	handlers []handler.Handler,
	renderer echo.Renderer,
) *echo.Echo {
	e := echo.New()
	e.Renderer = renderer

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.RequestID())

	for _, m := range middlewares {
		e.Use(m.Handler())
	}

	groups := map[string]*echo.Group{}
	for _, h := range handlers {
		group, ok := groups[h.Group()]
		if !ok {
			group = e.Group(h.Group())
			groups[h.Group()] = group
		}
		group.Add(h.Method(), h.Path(), h.Handle)
	}

	return e
}

func startServer(lc fx.Lifecycle, e *echo.Echo, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				addr := fmt.Sprintf(":%d", cfg.WebServerPort)
				if err := e.Start(addr); err != http.ErrServerClosed {
					e.Logger.Fatal("failed to start server")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Server.Shutdown(ctx)
		},
	})
}

func newRenderer(cfg *config.Config) echo.Renderer {
	render := pongo2echo.NewRenderer()
	render.AddDirectory(cfg.TemplatesDir)
	return render
}
