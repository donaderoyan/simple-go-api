package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID       string `gorm:"primaryKey;not null;unique"`
	ParentID string `gorm:"index"`
	User     User
	UserID   string `gorm:"index;not null"`

	Sku              string          `gorm:"type:varchar(100);index"`
	Name             string          `gorm:"type:varchar(255);not null"`
	Slug             string          `gorm:"type:varchar(255);not null"`
	Price            decimal.Decimal `gorm:"type:decimal(16,2);"`
	Stock            int
	Weight           decimal.Decimal `gorm:"type:decimal(10,2);"`
	ShortDescription string          `gorm:"type:text"`
	Description      string          `gorm:"type:text"`
	Status           int             `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt

	ProductImages []ProductImage
	Categories    []Category `gorm:"many2many:product_categories;"`
	CategoryID    []string   `json:"category_id" gorm:"-"`
}

func (entity *Product) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Product) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (p *Product) GetProducts(db *gorm.DB, perPage int, page int) (*[]Product, int64, error) {
	var err error
	var products []Product
	var count int64

	err = db.Debug().Model(&Product{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	err = db.Debug().Model(&Product{}).Order("created_at desc").Limit(perPage).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return &products, count, nil
}

func (p *Product) FindBySlug(db *gorm.DB, slug string) (*Product, error) {
	var err error
	var product Product

	err = db.Debug().Preload("ProductImage").Model(&Product{}).Where("slug = ?", slug).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Product) FindByID(db *gorm.DB, productID string) (*Product, error) {
	var err error
	var product Product

	err = db.Debug().Preload("ProductImage").Model(&Product{}).Where("id = ?", productID).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}
