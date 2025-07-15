package main

import (
	"os"

	"github.com/agungpg/perdana-task-manager/config"
	"github.com/agungpg/perdana-task-manager/internal/auth"
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

	// taskRepo := task.NewRepository(database.DB)
	// taskService := task.NewService(taskRepo)
	// taskHandler := task.NewHandler(taskService)
	// taskHandler.RegisterRoutes(app)

	// api := app.Group("/api", auth.JWTMiddleware())

	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
