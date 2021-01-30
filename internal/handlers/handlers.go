package handlers

import (
	"fmt"
	"github.com/gofiber/fiber"
	"net/http"
	"path/filepath"
	"rest-api/database"
	db "sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/models"
)

func GetCategories(ctx *fiber.Ctx) {
	var categories []models.Category
	database.DB.Find(&categories)
	ctx.JSON(categories)
}

func CreateCategory(ctx *fiber.Ctx) {
	category := new(models.Category)

	if err := ctx.BodyParser(category); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(err)
		return
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(err)
	}

	relativeFilePath := fmt.Sprintf(
		"%s%s%s",
		models.FilePath,
		string(filepath.Separator),
		image.Filename,
	)
	fullPath, err := filepath.Abs(fmt.Sprintf(
		"../assets/%s",
		relativeFilePath,
	))
	if err != nil {
		ctx.JSON(err)
	}

	err = ctx.SaveFile(image, fullPath)
	if err != nil {
		ctx.JSON(err)
	}

	category.File = models.File{
		Path: relativeFilePath,
	}
	errors := category.Validate()
	if errors != nil {
		ctx.JSON(fiber.Map{
			"success": false,
			"errors":  errors,
		})
		return
	}

	db.DBConn.Create(&category)

	ctx.JSON(fiber.Map{
		"success": true,
		"message": "Категория успешно создана",
	})
}
