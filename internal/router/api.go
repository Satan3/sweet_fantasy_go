package router

import (
	"github.com/gofiber/fiber"
	"sweet_fantasy_go/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/files", "../assets")

	group := app.Group("/categories")
	group.Get("/list", handlers.GetCategories)
	group.Post("/create", handlers.CreateCategory)
	group.Get("/delete/:id", handlers.DeleteCategory)
}
