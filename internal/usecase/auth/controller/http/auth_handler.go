package http

import (
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	config.Configuration
	domain.AuthService
}

func NewAuthHandler(authService domain.AuthService, config *config.Configuration) AuthHandler {
	return AuthHandler{
		AuthService:   authService,
		Configuration: *config,
	}
}

func (h AuthHandler) InitRoute(app fiber.Router) {
	app.Get("/role-permissions", h.GetUserRolePermissions)
}

func (h AuthHandler) GetUserRolePermissions(c *fiber.Ctx) error {
	ctx := c.UserContext()
	results, err := h.AuthService.GetUserRolePermissions(ctx, "bc9e0943-be4d-4735-bdf5-fbeffa59a115")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":   results,
		"status": fiber.StatusOK,
	})
}
