package handlers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
	"net/http"
	db "sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/models"
	"sweet_fantasy_go/internal/services"
	"sweet_fantasy_go/internal/validation"
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

	file, err := services.CreateAndSaveFile(image, models.FilePath)
	if err != nil {
		ctx.JSON(err)
		return
	}

	category.File = *file
	validationErrors := validation.Validate(category)
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

func UpdateCategory(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	category := new(models.Category)
	result := db.DBConn.Joins("File").First(&category, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Не существует категории с Id: %s", id),
		})
	}

	if err := ctx.BodyParser(category); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(err)
		return
	}

	image, _ := ctx.FormFile("image")
	if image != nil {
		err := services.ReplaceFile(image, models.FilePath, &category.File)
		if err != nil {
			ctx.JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	}
	validationErrors := validation.Validate(category)
	if len(validationErrors) >= 1 {
		ctx.JSON(fiber.Map{
			"success": false,
			"errors":  validationErrors,
		})
		return
	}
	db.DBConn.Save(&category)
	ctx.JSON(fiber.Map{
		"success": true,
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
