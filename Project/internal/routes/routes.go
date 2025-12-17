package routes

import (
	"project/internal/handler"
	"project/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	api := app.Group("/")

	api.Post("/users", userHandler.CreateUser)
	api.Get("/users/:id", userHandler.GetUserByID)
	api.Get("/users", userHandler.ListUsers)
	api.Put("/users/:id", userHandler.UpdateUser)
	api.Delete("/users/:id", userHandler.DeleteUser)
}
