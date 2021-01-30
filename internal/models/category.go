package models

import (
	"github.com/go-playground/validator"
	"gorm.io/gorm"
	"sweet_fantasy_go/internal/validation"
)

const FilePath = "categories"

type Category struct {
	gorm.Model
	Name        string `validate:"required" json:"name"`
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	Keywords    string `validate:"required" json:"keywords"`
	FileId      uint
	File        File
}

func (category *Category) Validate() []*validation.Error {
	var errors []*validation.Error
	validate := validator.New()
	err := validate.Struct(category)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := validation.Error{
				FailedField: err.Field(),
				Tag:         err.Tag(),
			}
			errors = append(errors, &element)
		}
	}
	return errors
}

func (category *Category) BeforeDelete(tx *gorm.DB) error {
	if err := category.File.removeFromStorage(); err != nil {
		return err
	}
	return nil
}
