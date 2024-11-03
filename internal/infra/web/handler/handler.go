package handler

import "github.com/labstack/echo/v4"

type Handler interface {
	Method() string
	Path() string
	Group() string
	Handle(c echo.Context) error
}

