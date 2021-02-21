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

	FileId uint `json:"-"`
	File   File `json:"file" gorm:"constraint:OnDelete:SET NULL"`
}

func (product *Product) BeforeDelete(db *gorm.DB) error {
	if err := product.File.RemoveFromStorage(); err != nil {
		return err
	}
	db.Delete(&product.File)
	return nil
}
