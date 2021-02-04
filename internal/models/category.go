package models

import (
	"gorm.io/gorm"
)

const FilePath = "categories"

type Category struct {
	Base
	Name        string `validate:"required" json:"name"`
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	Keywords    string `validate:"required" json:"keywords"`
	FileId      uint   `json:"-"`
	File        File   `json:"file" gorm:"constraint:OnDelete:SET NULL"`
}

func (category *Category) BeforeDelete(db *gorm.DB) error {
	if err := category.File.RemoveFromStorage(); err != nil {
		return err
	}
	db.Delete(&category.File)
	return nil
}
