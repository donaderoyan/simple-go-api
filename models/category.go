package model

import "time"

type Category struct {
	ID        string `gorm:"primaryKey;unique;not null"`
	ParentID  string `gorm:"index"`
	Section   Section
	SectionID string    `gorm:"type:varchar(36);index"`
	Products  []Product `gorm:"many2many:product_categories;"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Slug      string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
