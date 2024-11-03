package user

import (
	"go-echo/internal/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type GetUserHandler struct {
	storage map[int]*domain.User
	logger  *zap.Logger
}

func NewGetUserHandler(logger *zap.Logger) *GetUserHandler {
	return &GetUserHandler{
		storage: map[int]*domain.User{
			1: {ID: 1, Name: "Alice", Email: "alice@gmail.com", Role: "admin"},
			2: {ID: 2, Name: "Bob", Email: "bob@gmail.com", Role: "user"},
		},
		logger: logger,
	}
}

func (h *GetUserHandler) Method() string {
	return http.MethodGet
}

func (h *GetUserHandler) Path() string {
	return "/:id"
}

func (h *GetUserHandler) Group() string {
	return "user"
}

func (h *GetUserHandler) Handle(c echo.Context) error {
	// User ID from path `users/:id`
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid user ID")
	}
	user := h.storage[id]

	err = c.Render(http.StatusOK, "pages/user/profile.html", map[string]any{
		"user": user,
	})
	if err != nil {
		h.logger.Error("Failed to render template", zap.Error(err))
		return c.String(http.StatusInternalServerError, "Internal server error")
	}
	return err
}
