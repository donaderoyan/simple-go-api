package model

import (
	"time"
)

type ProductImage struct {
	ID         string `gorm:"primaryKey;not null;unique"`
	Product    Product
	ProductID  string `gorm:"index;not null"`
	Path       string `gorm:"type:text"`
	ExtraLarge string `gorm:"type:text"`
	Large      string `gorm:"type:text"`
	Medium     string `gorm:"type:text"`
	Small      string `gorm:"type:text"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
