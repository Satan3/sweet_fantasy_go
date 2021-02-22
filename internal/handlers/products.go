package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"sweet_fantasy_go/internal/models"
	categoryRepository "sweet_fantasy_go/internal/repositories/categories_repository"
	filesRepository "sweet_fantasy_go/internal/repositories/files_repository"
	productRepository "sweet_fantasy_go/internal/repositories/product_repository"
	"sweet_fantasy_go/internal/validation"
)

func GetProducts(ctx *fiber.Ctx) error {
	return successResponse(ctx, productRepository.FindAll())
}

func CreateProduct(ctx *fiber.Ctx) error {
	product := new(models.Product)

	if err := ctx.BodyParser(product); err != nil {
		return errorResponse(ctx.Status(http.StatusBadRequest), err.Error())
	}

	if product.FileId == 0 {
		return errorResponse(ctx, "Отсутствует идентификатор файла обложки")
	}

	stringFileId := strconv.FormatUint(uint64(product.FileId), 10)
	file, err := filesRepository.FindById(stringFileId)
	if err != nil {
		return errorResponse(ctx, "Файл обложки не найден")
	}

	if product.CategoryId == 0 {
		return errorResponse(ctx, "Отсутствует идентификатор категории")
	}

	stringCategoryId := strconv.FormatUint(uint64(product.CategoryId), 10)
	category, err := categoryRepository.FindById(stringCategoryId)
	if err != nil {
		return errorResponse(ctx, fmt.Sprintf("Отсутствует категория с id: %s", stringCategoryId))
	}

	product.Category = *category
	product.File = *file
	fmt.Println(product)
	validationErrors := validation.Validate(product)
	if len(validationErrors) >= 1 {
		return errorResponse(ctx, fiber.Map{"errors": validationErrors})
	}

	productRepository.Create(product)
	return successResponse(ctx, "Товар успешно создан")
}
