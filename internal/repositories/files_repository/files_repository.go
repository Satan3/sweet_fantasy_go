package files_repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	db "sweet_fantasy_go/internal/database"
	"sweet_fantasy_go/internal/models"
)

func FindById(id string) (*models.File, error) {
	file := new(models.File)
	result := db.DBConn.First(&file, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(fmt.Sprintf("Не существует категории с Id: %s", id))
	}
	return file, nil
}

func Create(file *models.File) {
	db.DBConn.Create(&file)
}

func Delete(file *models.File) {
	db.DBConn.Delete(&file)
}
