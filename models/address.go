package model

import "time"

type Address struct {
	ID         string `gorm:"primaryKey;not null"`
	User       User
	UserID     string `gorm:"type:varchar(255);not null;unique;index"`
	Name       string `gorm:"type:varchar(255);not null"`
	IsPrimary  bool
	CityID     string `gorm:"type:varchar(255);not null"`
	ProvinceID string `gorm:"type:varchar(255);not null"`
	Address1   string `gorm:"type:varchar(255);not null"`
	Address2   string `gorm:"type:varchar(255);not null"`
	Phone      string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:varchar(255);not null"`
	PostCode   string `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
