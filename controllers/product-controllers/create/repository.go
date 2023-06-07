package createProduct

import (
	model "github.com/donaderoyan/simple-go-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateProductRepository(input *model.Product) (*model.Product, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateProductRepository(input *model.Product) (*model.Product, string) {
	var product model.Product
	db := r.db.Model(&product)
	errorCode := make(chan string, 1)

	var user model.User
	checkUserAccount := r.db.Model(&user).Debug().Select("*").Where("id = ?", input.ID).Find(&user)
	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "CREATE_PRODUCT_CONFLICT_409"
		return &product, <-errorCode
	}

	addNewProduct := db.Debug().Create(input)
	db.Commit()

	if addNewProduct.Error != nil {
		errorCode <- "CREATE_PRODUCT_FAILED_403"
		return input, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return input, <-errorCode
}
