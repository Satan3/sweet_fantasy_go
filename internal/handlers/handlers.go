package handlers

import (
	"fmt"
	"github.com/gofiber/fiber"
	"net/http"
	"rest-api/database"
	"sweet_fantasy_go/internal/models"
)

func GetCategories(ctx *fiber.Ctx) {
	var categories []models.Category
	database.DB.Find(&categories)
	ctx.JSON(categories)
}

func CreateCategory(ctx *fiber.Ctx) {
	category := new(models.Category)

	fmt.Println(category)

	if err := ctx.BodyParser(category); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(err)
		return
	}

	errors := category.Validate()
	fmt.Println("errors", errors)

	if errors != nil {
		ctx.JSON(fiber.Map{
			"success": false,
			"errors":  errors,
		})
		return
	}

	ctx.JSON(ctx.JSON(fiber.Map{
		"success": true,
		"message": "Категория успешно создана",
	}))
}
