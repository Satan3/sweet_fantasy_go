package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"sweet_fantasy_go/internal/models"
	categoriesRepository "sweet_fantasy_go/internal/repositories/categories_repository"
	filesRepository "sweet_fantasy_go/internal/repositories/files_repository"
	"sweet_fantasy_go/internal/validation"
)

func GetCategories(ctx *fiber.Ctx) error {
	return successResponse(ctx, categoriesRepository.FindAll())
}

func CreateCategory(ctx *fiber.Ctx) error {
	category := new(models.Category)

	if err := ctx.BodyParser(category); err != nil {
		return errorResponse(ctx.Status(http.StatusBadRequest), err.Error())
	}

	if category.FileId == 0 {
		return errorResponse(ctx, "Отсутствует идентификатор файла обложки")
	}

	stringFileId := strconv.FormatUint(uint64(category.FileId), 10)
	file, err := filesRepository.FindById(stringFileId)
	if err != nil {
		return errorResponse(ctx, "Файл обложки не найден")
	}

	category.File = *file
	validationErrors := validation.Validate(category)
	if len(validationErrors) >= 1 {
		return errorResponse(ctx, fiber.Map{"errors": validationErrors})
	}

	categoriesRepository.Create(category)
	return successResponse(ctx, "Категория успешно создана")
}

func UpdateCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	category, err := categoriesRepository.FindById(id)
	if err != nil {
		return errorResponse(ctx, "Не найдена категория для обновления")
	}
	prevFileId := category.FileId

	if err := ctx.BodyParser(category); err != nil {
		return errorResponse(ctx.Status(http.StatusBadRequest), "Неверная структура запроса")
	}

	if category.FileId == 0 {
		return errorResponse(ctx, "Отсутствует идентификатор файла обложки")
	}

	if prevFileId != category.FileId {
		stringFileId := strconv.FormatUint(uint64(category.FileId), 10)
		file, err := filesRepository.FindById(stringFileId)
		if err != nil {
			return errorResponse(ctx, "Файл обложки не найден")
		}

		filesRepository.Delete(&category.File)
		category.File = *file
	}

	validationErrors := validation.Validate(category)
	if len(validationErrors) >= 1 {
		return errorResponse(ctx, fiber.Map{"errors": validationErrors})
	}

	categoriesRepository.Update(category)
	return successResponse(ctx, "Категория успешно обновлена")
}

func DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	category, err := categoriesRepository.FindById(id)
	if err != nil {
		return errorResponse(ctx, err.Error())
	}

	categoriesRepository.Delete(category)
	return successResponse(ctx, "Категория успешно удалена")
}
