package router

import (
	"github.com/gofiber/fiber/v2"
	"sweet_fantasy_go/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/static", "./assets")

	fileGroup := app.Group("/files")
	fileGroup.Post("/upload", handlers.Upload)

	categoryGroup := app.Group("/categories")
	categoryGroup.Post("/list", handlers.GetCategories)
	categoryGroup.Post("/create", handlers.CreateCategory)
	categoryGroup.Get("/get/:id", handlers.GetCategory)
	categoryGroup.Put("/update/:id", handlers.UpdateCategory)
	categoryGroup.Delete("/delete/:id", handlers.DeleteCategory)

	productGroup := app.Group("/products")
	productGroup.Get("/list", handlers.GetProducts)
	productGroup.Post("/create", handlers.CreateProduct)
}
