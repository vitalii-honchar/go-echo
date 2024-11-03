package middleware

import "github.com/labstack/echo/v4"

type Middleware interface {
	Handler() echo.MiddlewareFunc
}
