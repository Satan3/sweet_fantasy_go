package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"os"
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

	router.SetupRoutes(app)
	database.InitDatabase()

	app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
