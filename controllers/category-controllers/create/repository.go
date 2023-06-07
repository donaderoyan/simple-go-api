package createCategory

import (
	model "github.com/donaderoyan/simple-go-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCategoryRepository(input *model.Category) (*model.Category, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateCategoryRepository(input *model.Category) (*model.Category, string) {
	var category model.Category
	db := r.db.Model(&category)
	errorCode := make(chan string, 1)

	addNewCategory := db.Debug().Create(input)
	db.Commit()

	if addNewCategory.Error != nil {
		errorCode <- "CREATE_CATEGORY_FAILED_403"
		return input, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return input, <-errorCode
}
