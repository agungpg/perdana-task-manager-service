package task

import (
	"fmt"
	"time"

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
	router.Post("/", h.CreateTask)
	router.Get("/", h.GetTaskList)
	router.Get("/:id", h.GetTaskByID)
}

func (h *Handler) CreateTask(c *fiber.Ctx) error {
	var input CreateTaskRequest
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	claims := c.Locals("userClaims").(jwt.MapClaims)
	userID := claims["id"].(string)

	// Parse input.DueDate (string) to time.Time
	var dueDate time.Time
	var err error
	if input.DueDate != "" {
		dueDate, err = time.Parse(time.RFC3339, input.DueDate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid due date format. Use RFC3339 format.",
			})
		}
	}

	err = h.service.CreateTask(
		c.Context(),
		input.ProjectID,
		input.StatusID,
		input.AssigneeID,
		input.ReporterID,
		input.Priority,
		input.Name,
		input.Description,
		input.Thumbnail,
		dueDate,
		userID,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
	})
}

func (h *Handler) GetTaskList(c *fiber.Ctx) error {
	projectId := c.Query("project_id")
	fmt.Printf("projectId: %s\n", projectId)
	tasks, err := h.service.GetTaskList(c.Context(), projectId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": tasks,
	})
}

func (h *Handler) GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := h.service.GetTaskByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(task)
}
