package util

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		limit, _ := c.MustGet("Limit").(int64)
// 		page, _ := c.MustGet("Page").(int64)

// 		offset := (page - 1) * limit

//			return db.Offset(int(offset)).Limit(int(limit))
//		}
//	}
type PaginateScope interface {
	PaginatedResult() *gorm.DB
}

type Paginate struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}

func NewPaginate(limit int, page int) *Paginate {
	return &Paginate{Limit: limit, Page: page}
}

func (p *Paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.Page - 1) * p.Limit

	logrus.Info("REPOSITORY pagination.go  ", offset, p)
	return db.Offset(offset).
		Limit(p.Limit)
}
