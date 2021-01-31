package main

import (
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/router"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic("There is no .env file")
	}
}

func main() {
	app := fiber.New()
	router.SetupRoutes(app)
	database.InitDatabase()

	app.Listen(":3000")
}
