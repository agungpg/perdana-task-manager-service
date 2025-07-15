package project

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/", h.CreateProduct)
	router.Get("/", h.GetProjectList)
	router.Get("/:id", h.GetProjectByID)
}

type createProductInput struct {
	Name        string `json:"name"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	Description string `json:"description,omitempty"`
}

func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	var input createProductInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}
	claims := c.Locals("userClaims").(jwt.MapClaims)
	userID := claims["id"].(string)
	err := h.service.CreateProject(c.Context(), userID, input.Name, input.Thumbnail, input.Description)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Project successfully created!",
	})
}

func (h *Handler) GetProjectList(c *fiber.Ctx) error {
	pageStr := c.Query("page", "1")          // default to "1" if not provided
	pageSizeStr := c.Query("pageSize", "10") // default to "10" if not provided
	name := c.Query("name", "")              // default to "10" if not provided

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	projects, total, err := h.service.GetProjectList(c.Context(), page, pageSize, name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"data": projects,
		"pagination": fiber.Map{
			"page":       page,
			"page_size":  pageSize,
			"total_data": total,
		},
	})
}

func (h *Handler) GetProjectByID(c *fiber.Ctx) error {
	id := c.Params("id") // Get the dynamic id from the path
	project, err := h.service.GetProjectByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Project not found"})
	}
	return c.JSON(project)
}
