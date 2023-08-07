package resultsProduct

import (
	model "github.com/donaderoyan/simple-go-api/models"
	util "github.com/donaderoyan/simple-go-api/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	ResultsProductRepository() (*[]model.Product, string)
	SetScopeResult(paginate util.Paginate) *repository
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) SetScopeResult(paginate util.Paginate) *repository {
	db := r.db
	var scopePaginate util.Paginate
	scopePaginate.Limit = paginate.Limit
	scopePaginate.Page = paginate.Page

	scope := db.Debug().Scopes(scopePaginate.PaginatedResult)
	return &repository{db: scope}
}

func (r *repository) ResultsProductRepository() (*[]model.Product, string) {

	var products []model.Product

	db := r.db.Model(&products)
	errorCode := make(chan string, 1)

	// resultProducts := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
	// 	return db.Select("ID", "FirstName", "LastName", "Email")
	// }).Select("*").Find(&products)
	resultProducts := db.Debug().Preload(clause.Associations).Select("*").Find(&products)

	if resultProducts.Error != nil {
		errorCode <- "RESULTS_STUDENT_NOT_FOUND_404"
		return &products, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &products, <-errorCode
}
