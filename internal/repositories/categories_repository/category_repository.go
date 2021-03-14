package categories_repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	db "sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/models"
)

type Pagination struct {
	Total int64 `json:"total"`
	Items []models.Category
}

func FindList(page int, pageSize int) (categories []*models.Category, total int64) {
	db.DBConn.
		Table("categories").
		Count(&total).
		Joins("File").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&categories)
	return
}

func FindById(id string) (*models.Category, error) {
	category := new(models.Category)
	result := db.DBConn.Joins("File").First(&category, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(fmt.Sprintf("Не существует категории с Id: %s", id))
	}
	return category, nil
}

func Create(category *models.Category) {
	db.DBConn.Create(&category)
}

func Update(category *models.Category) {
	db.DBConn.Save(&category)
}

func Delete(category *models.Category) {
	db.DBConn.Delete(&category)
}
