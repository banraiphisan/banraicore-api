package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tubfuzzy/banraicore-api/config"
	"github.com/tubfuzzy/banraicore-api/internal/domain"
)

type ShortUrlHandler struct {
	config.Configuration
	domain.ShortUrlService
}

func NewShortUrlHandler(shortUrlService domain.ShortUrlService, config *config.Configuration) ShortUrlHandler {
	return ShortUrlHandler{
		ShortUrlService: shortUrlService,
		Configuration:   *config,
	}
}

func (h ShortUrlHandler) InitRoute(app fiber.Router) {
	app.Get("/r/:code", h.ShortenRedirect)
	app.Post("/shorten", h.CreateShorten)
	app.Get("/shorten", h.GetShortenList)
	app.Get("/shorten/:code", h.GetShortenByCode)
	app.Delete("/shorten/:code", h.DeleteShortenByCode)
}

func (h ShortUrlHandler) ShortenRedirect(c *fiber.Ctx) error {
	code := c.Params("code")
	result, err := h.ShortUrlService.GetByCode(c.Context(), code)
	if err != nil || result == nil {
		return c.Status(fiber.StatusNotFound).SendString("URL not found")
	}
	return c.Redirect(result.TargetUrl, fiber.StatusFound)
}

func (h ShortUrlHandler) CreateShorten(c *fiber.Ctx) error {
	var payload domain.CreateShortUrlPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	result, err := h.ShortUrlService.Create(c.Context(), payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h ShortUrlHandler) GetShortenList(c *fiber.Ctx) error {
	list, err := h.ShortUrlService.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(list)
}

func (h ShortUrlHandler) GetShortenByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	result, err := h.ShortUrlService.GetByCode(c.Context(), code)
	if err != nil || result == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "not found",
		})
	}
	return c.JSON(result)
}

func (h ShortUrlHandler) DeleteShortenByCode(c *fiber.Ctx) error {
	code := c.Params("code")
	err := h.ShortUrlService.DeleteByCode(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
