package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/router"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic("There is no .env file")
	}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.InitDatabase()
	router.SetupRoutes(app)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
