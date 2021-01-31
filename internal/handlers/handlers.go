package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	db "sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/models"
)

func GetCategories(ctx *fiber.Ctx) {
	var categories []models.Category
	db.DBConn.Joins("File").Find(&categories)
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
		ctx.JSON(fiber.Map{
			"success": false,
			"message": "Отсутствует файл обложки",
		})
		return
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
		return
	}

	err = ctx.SaveFile(image, fullPath)
	if err != nil {
		ctx.JSON(err)
		return
	}

	category.File = models.File{
		Path: relativeFilePath,
	}
	validationErrors := category.Validate()
	if validationErrors != nil {
		ctx.JSON(fiber.Map{
			"success":          false,
			"validationErrors": validationErrors,
		})
		return
	}

	db.DBConn.Create(&category)

	ctx.JSON(fiber.Map{
		"success": true,
		"message": "Категория успешно создана",
	})
}

func DeleteCategory(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	category := new(models.Category)
	result := db.DBConn.Joins("File").First(&category, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Не существует категории с Id: %s", id),
		})
		return
	}
	db.DBConn.Delete(&category)
}
