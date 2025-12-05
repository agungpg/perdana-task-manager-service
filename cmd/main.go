package main

import (
	"fmt"
	"log"
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

	// Run migrations when invoked with: ./main migrate [up|down]
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		direction := "up"
		if len(os.Args) > 2 {
			direction = os.Args[2]
		}

		if err := database.RunMigrations(direction); err != nil {
			log.Fatalf("Migration %s failed: %v", direction, err)
		}

		fmt.Printf("Migration %s completed successfully\n", direction)
		return
	}

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
	if port == "" {
		port = "8080"
	}
	if err := app.Listen(":" + port); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
