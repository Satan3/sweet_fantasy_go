package models

import "gorm.io/gorm"

const ProductPath = "products"

type Product struct {
	Base
	Name        string `validate:"required" json:"name"`
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
	Keywords    string `validate:"required" json:"keywords"`
	Price       int    `validate:"required,numeric" json:"price"`
	Discount    int    `json:"discount"`
	HitSales    bool   `json:"hit_sales"`

	CategoryId uint     `json:"category_id"`
	Category   Category `json:"category" gorm:"constraint:OnDelete:SET NULL"`

	FileId uint `validate:"required" json:"file_id"`
	File   File `validate:"required" json:"file" gorm:"constraint:OnDelete:SET NULL"`
}

func (product *Product) BeforeDelete(db *gorm.DB) error {
	product.File.RemoveFromStorage()
	db.Delete(&product.File)
	return nil
}
