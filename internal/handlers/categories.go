package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"sweet_fantasy_go/internal/models"
	categoriesRepository "sweet_fantasy_go/internal/repositories/categories_repository"
	"sweet_fantasy_go/internal/services"
	"sweet_fantasy_go/internal/validation"
)

func GetCategories(ctx *fiber.Ctx) error {
	return ctx.JSON(categoriesRepository.FindAll())
}

func CreateCategory(ctx *fiber.Ctx) error {
	category := new(models.Category)

	if err := ctx.BodyParser(category); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	image, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(fiber.Map{
			"success": false,
			"message": "Отсутствует файл обложки",
		})
	}

	file, err := services.CreateAndSaveFile(image, models.CategoryPath)
	if err != nil {
		return ctx.JSON(err)
	}

	category.File = *file
	validationErrors := validation.Validate(category)
	if len(validationErrors) >= 1 {
		return ctx.JSON(fiber.Map{
			"success": false,
			"errors":  validationErrors,
		})
	}

	categoriesRepository.Create(category)
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Категория успешно создана",
	})
}

func UpdateCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	category, err := categoriesRepository.FindById(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := ctx.BodyParser(category); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	image, _ := ctx.FormFile("image")
	if image != nil {
		err := services.ReplaceFile(image, models.CategoryPath, &category.File)
		if err != nil {
			return ctx.JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
	}

	validationErrors := validation.Validate(category)
	if len(validationErrors) >= 1 {
		return ctx.JSON(fiber.Map{
			"success": false,
			"errors":  validationErrors,
		})
	}

	categoriesRepository.Update(category)
	return ctx.JSON(fiber.Map{
		"success": true,
	})
}

func DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	category, err := categoriesRepository.FindById(id)
	if err != nil {
		return ctx.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	categoriesRepository.Delete(category)
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Категория успешно удалена",
	})
}
