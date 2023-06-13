package model

import (
	"time"

	util "github.com/donaderoyan/simple-go-api/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"primaryKey;not null"`
	Addresses     []Address
	FirstName     string    `gorm:"type:varchar(255);not null"`
	LastName      string    `gorm:"type:varchar(255);not null"`
	Email         string    `gorm:"type:varchar(255);unique;not null"`
	Password      string    `gorm:"type:varchar(255);not null" json:"-"`
	RememberToken string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

func (entity *User) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *User) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
