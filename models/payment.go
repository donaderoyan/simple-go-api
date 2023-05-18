package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct {
	ID          string `gorm:"primaryKey;not null;unique"`
	Order       Order
	OrderID     string          `gorm:"index;not null"`
	Number      string          `gorm:"type:varchar(100);index"`
	Amount      decimal.Decimal `gorm:"type:decimal(16,2)"`
	Method      string          `gorm:"type:varchar(100)"`
	Status      string          `gorm:"type:varchar(100)"`
	Token       string          `gorm:"type:varchar(100);index"`
	Payload     string          `gorm:"type:text"`
	PaymentType string          `gorm:"type:varchar(100)"`
	VaNumber    string          `gorm:"type:varchar(100)"`
	BillCode    string          `gorm:"type:varchar(100)"`
	BillKey     string          `gorm:"type:varchar(100)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
