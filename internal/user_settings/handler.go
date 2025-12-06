package task

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Get("/", h.GetUserSettings)
}

func (h *Handler) GetUserSettings(c *fiber.Ctx) error {
	claims := c.Locals("userClaims").(jwt.MapClaims)
	userID := claims["id"].(string)
	userSettings, err := h.service.GetUserSettings(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(userSettings)
}
