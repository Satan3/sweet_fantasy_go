package handlers

import (
	"github.com/gofiber/fiber/v2"
	"sweet_fantasy_go/internal/services"
)

func Upload(ctx *fiber.Ctx) error {
	storagePath := ctx.FormValue("storage_path")
	if storagePath == "" {
		return errorResponse(ctx, "Не указано хранилище файла")
	}

	formFile, err := ctx.FormFile("file")
	if err != nil {
		return errorResponse(ctx, "Файл отсутствует")
	}

	file, err := services.CreateAndSaveFile(formFile, storagePath)
	if err != nil {
		return errorResponse(ctx, err.Error())
	}

	return successResponse(ctx, file.ID)
}

func successResponse(ctx *fiber.Ctx, payload interface{}) error {
	return ctx.JSON(fiber.Map{
		"success": true,
		"payload": payload,
	})
}

func errorResponse(ctx *fiber.Ctx, payload interface{}) error {
	return ctx.JSON(fiber.Map{
		"success": false,
		"payload": payload,
	})
}
