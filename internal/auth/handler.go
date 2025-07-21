package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// Change parameter type to fiber.Router to match project handler pattern
func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/register", h.Register)
	router.Post("/login", h.Login)
}

type authInput struct {
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var input authInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	fmt.Println("Registering user:", input.Username)
	fmt.Println("Registering email:", input.Email)
	fmt.Println("Registering password:", input.Password)
	err := h.service.Register(c.Context(), input.Username, input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var input authInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	token, err := h.service.Login(c.Context(), input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
