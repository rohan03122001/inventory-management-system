package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:255;not null" json:"name" binding:"required"`
	Description string         `gorm:"type:text" json:"description"`
	Price       float64         `gorm:"not null" json:"price" binding:"required,gt=0"`
	Quantity    int            `gorm:"not null" json:"quantity" binding:"required,gte=0"`
	SKU         string         `gorm:"size:50;unique;not null" json:"sku" binding:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}


func (p *Product)BeforeCreate(tx *gorm.DB) error{
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

func(p *Product)BeforeUpdate(tx *gorm.DB) error{
	p.UpdatedAt= time.Now()
	return nil
}