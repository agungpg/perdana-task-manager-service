package main

import (
	"os"

	"github.com/agungpg/perdana-task-manager/config"
	"github.com/agungpg/perdana-task-manager/internal/auth"
	"github.com/agungpg/perdana-task-manager/internal/project"
	"github.com/agungpg/perdana-task-manager/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	database.Connect()

	app := fiber.New()

	// Auth setup
	authRepo := auth.NewRepository(database.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(app)

	api := app.Group("/api") // No middleware here

	projectRepo := project.NewRepository(database.DB)
	projectService := project.NewService(projectRepo)
	projectHandler := project.NewHandler(projectService)

	// Apply JWTMiddleware only to /project routes
	projectGroup := api.Group("/project", auth.JWTMiddleware())
	projectHandler.RegisterRoutes(projectGroup)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
