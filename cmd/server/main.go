package main

import (
	"net/http"
	"strconv"

	"github.com/h3poteto/pongo2echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func main() {
	e := echo.New()
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	render := pongo2echo.NewRenderer()
	render.AddDirectory("../../web/templates")

	e.Renderer = render

	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(loggingMiddleware(log))

	storage := map[int]*User{
		1: {ID: 1, Name: "Alice", Email: "alice@gmail.com", Role: "admin"},
		2: {ID: 2, Name: "Bob", Email: "bob@gmail.com", Role: "user"},
	}

	

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser(storage, log))

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(storage map[int]*User, log *zap.Logger) echo.HandlerFunc {
	return func(c echo.Context) error {
		// User ID from path `users/:id`
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid user ID")
		}
		user := storage[id]

		err = c.Render(http.StatusOK, "pages/user/profile.html", map[string]any{
			"user": user,
		})
		if err != nil {
			log.Error("Failed to render template", zap.Error(err))
			return c.String(http.StatusInternalServerError, "Internal server error")
		}
		return err
	}
}

func loggingMiddleware(log *zap.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
				zap.String("method", c.Request().Method),
				zap.String("requestID", c.Response().Header().Get(echo.HeaderXRequestID)),
			)

			return nil
		},
	})
}
