package handlers

import (
	"github.com/gofiber/fiber"
	"net/http"
	"sweet_fantasy_go/internal/models"
	categoriesRepository "sweet_fantasy_go/internal/repositories/categories_repository"
	"sweet_fantasy_go/internal/services"
	"sweet_fantasy_go/internal/validation"
)

func GetCategories(ctx *fiber.Ctx) {
	ctx.JSON(categoriesRepository.FindAll())
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

	file, err := services.CreateAndSaveFile(image, models.CategoryPath)
	if err != nil {
		ctx.JSON(err)
		return
	}

	category.File = *file
	validationErrors := validation.Validate(category)
	if len(validationErrors) >= 1 {
		ctx.JSON(fiber.Map{
			"success": false,
			"errors":  validationErrors,
		})
		return
	}

	categoriesRepository.Create(category)
	ctx.JSON(fiber.Map{
		"success": true,
		"message": "Категория успешно создана",
	})
}

func UpdateCategory(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	category, err := categoriesRepository.FindById(id)
	if err != nil {
		ctx.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := ctx.BodyParser(category); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(err)
		return
	}

	image, _ := ctx.FormFile("image")
	if image != nil {
		err := services.ReplaceFile(image, models.CategoryPath, &category.File)
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
	categoriesRepository.Update(category)
	ctx.JSON(fiber.Map{
		"success": true,
	})
}

func DeleteCategory(ctx *fiber.Ctx) {
	id := ctx.Params("id")
	category, err := categoriesRepository.FindById(id)
	if err != nil {
		ctx.JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	categoriesRepository.Delete(category)
	ctx.JSON(fiber.Map{
		"success": true,
		"message": "Категория успешно удалена",
	})
}
