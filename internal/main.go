package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/handlers"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic("There is no .env file")
	}
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	database.InitDatabase()

	app.Listen(":3000")
}

func setupRoutes(app *fiber.App) {
	app.Static("/files", "../assets")

	group := app.Group("/categories")
	group.Get("/list", handlers.GetCategories)
	group.Post("/create", handlers.CreateCategory)
}
