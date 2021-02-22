package router

import (
	"github.com/gofiber/fiber/v2"
	"sweet_fantasy_go/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/static", "../assets")

	fileGroup := app.Group("/files")
	fileGroup.Post("/upload", handlers.Upload)

	group := app.Group("/categories")
	group.Get("/list", handlers.GetCategories)
	group.Post("/create", handlers.CreateCategory)
	group.Put("/update/:id", handlers.UpdateCategory)
	group.Delete("/delete/:id", handlers.DeleteCategory)
}
