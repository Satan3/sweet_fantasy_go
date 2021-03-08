package models

import (
	"gorm.io/gorm"
)

const CategoryPath = "categories"

type Category struct {
	Base
	Name        string `validate:"required" json:"name"`
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	Keywords    string `validate:"required" json:"keywords"`

	FileId uint `validate:"required" json:"file_id"`
	File   File `validate:"required" json:"file" gorm:"constraint:OnDelete:SET NULL"`
}

func (category *Category) BeforeDelete(db *gorm.DB) error {
	category.File.RemoveFromStorage()
	db.Delete(&category.File)
	return nil
}
