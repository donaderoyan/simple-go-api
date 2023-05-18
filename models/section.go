package model

import "time"

type Section struct {
	ID         string `gorm:"primaryKey;not null;unique"`
	Name       string `gorm:"type:varchar(255);not null"`
	Slug       string `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Categories []Category
}
