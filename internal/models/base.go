package models

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (base *Base) BeforeCreate(db *gorm.DB) error {
	now := time.Now()
	base.CreatedAt, base.UpdatedAt = now, now
	return nil
}

func (base *Base) BeforeUpdate(db *gorm.DB) error {
	base.UpdatedAt = time.Now()
	return nil
}
