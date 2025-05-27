package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/banraiphisan/banraicore-api/config"
	"github.com/banraiphisan/banraicore-api/internal/domain"
	"github.com/banraiphisan/banraicore-api/pkg/utils"
	"strconv"
)

type UserHandler struct {
	config.Configuration
	domain.UserService
}

func NewUserHandler(userService domain.UserService, config *config.Configuration) UserHandler {
	return UserHandler{
		UserService:   userService,
		Configuration: *config,
	}
}

func (h UserHandler) InitRoute(app fiber.Router) {
	api := app.Group("/users")
	api.Post("/create", h.CreateUser)
	api.Get("/", h.GetUsers)
	api.Get("/:id", h.GetUserByID)
	api.Put("/:id", h.UpdateUser)

}

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	var payload domain.CreateOrUpdateUserPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := utils.ValidateStruct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.UserService.CreateUser(c.Context(), domain.CreateOrUpdateUserPayload{
		Username: payload.Username,
		Email:    payload.Email,
		Password: payload.Password,
		RoleCode: payload.RoleCode,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (h UserHandler) GetUsers(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page number"})
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize", "10"))
	if err != nil || pageSize < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid page size"})
	}

	users, total, err := h.UserService.GetUsers(c.Context(), page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch users"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"users": users,
		"pagination": fiber.Map{
			"page":         page,
			"pageSize":     pageSize,
			"totalPages":   (total + int64(pageSize) - 1) / int64(pageSize),
			"totalRecords": total,
		},
	})
}

func (h UserHandler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")

	id, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	user, err := h.UserService.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h UserHandler) UpdateUser(c *fiber.Ctx) error {
	userIDParam := c.Params("id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var payload domain.CreateOrUpdateUserPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	payload.ID = &userID

	if err := utils.ValidateStruct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.UserService.UpdateUser(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
}
