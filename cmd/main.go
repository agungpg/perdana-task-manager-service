package main

import (
	"fmt"
	"os"

	"github.com/agungpg/perdana-task-manager/config"
	"github.com/agungpg/perdana-task-manager/internal/auth"
	"github.com/agungpg/perdana-task-manager/internal/project"
	"github.com/agungpg/perdana-task-manager/internal/task"
	"github.com/agungpg/perdana-task-manager/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	if err := database.Connect(); err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	app := fiber.New()

	// Auth setup
	authRepo := auth.NewRepository(database.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// Create auth group and pass it to RegisterRoutes
	authGroup := app.Group("/auth")
	authHandler.RegisterRoutes(authGroup)

	api := app.Group("/api") // No middleware here

	projectRepo := project.NewRepository(database.DB)
	projectService := project.NewService(projectRepo)
	projectHandler := project.NewHandler(projectService)

	// Apply JWTMiddleware only to /project routes
	projectGroup := api.Group("/project", auth.JWTMiddleware())
	projectHandler.RegisterRoutes(projectGroup)

	taskRepo := task.NewRepository(database.DB)
	taskService := task.NewService(taskRepo)
	taskHandler := task.NewHandler(taskService)
	taskGroup := api.Group("/task", auth.JWTMiddleware())
	taskHandler.RegisterRoutes(taskGroup)

	port := os.Getenv("PORT")
	if err := app.Listen(":" + port); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
