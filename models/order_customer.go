package model

import "time"

type OrderCustomer struct {
	ID         string `gorm:"primaryKey;not null;unique"`
	User       User
	UserID     string `gorm:"index;not null"`
	Order      Order
	OrderID    string `gorm:"index;not null"`
	FirstName  string `gorm:"type:varchar(100);not null"`
	LastName   string `gorm:"type:varchar(100);not null"`
	CityID     string `gorm:"type:varchar(100);"`
	ProvinceID string `gorm:"type:varchar(100);"`
	Address1   string `gorm:"type:varchar(100);"`
	Address2   string `gorm:"type:varchar(100);"`
	Phone      string `gorm:"type:varchar(50);"`
	Email      string `gorm:"type:varchar(100);"`
	PostCode   string `gorm:"type:varchar(100);"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
