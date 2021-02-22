package product_repository

import (
	db "sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/models"
)

func FindAll() (products []models.Product) {
	db.DBConn.Joins("File").Joins("Category").Find(&products)
	return
}

func Create(model *models.Product) {
	db.DBConn.Create(&model)
}

func Update(model *models.Product) {
	db.DBConn.Save(&model)
}

func Delete(model *models.Product) {
	db.DBConn.Delete(&model)
}
