package model

type Shipment struct {
	ID          string `gorm:"primaryKey;not null;unique"`
	User        User
	UserID      string `gorm:"index;not null"`
	Order       Order
	OrderID     string `gorm:"index;not null"`
	TrackNumber string `gorm:"type:varchar(255);index"`
	Status      string `gorm:"type:varchar(36);index"`
}
