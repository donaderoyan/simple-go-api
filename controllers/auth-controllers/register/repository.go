package registerAuth

import (
	model "github.com/donaderoyan/simple-go-api/models"
	"gorm.io/gorm"
)

type Repository interface {
	RegisterRepository(input *model.User) (*model.User, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) RegisterRepository(input *model.User) (*model.User, string) {
	var user model.User
	db := r.db.Model(&user)
	errorCode := make(chan string, 1)

	// err := db.Debug().Model(&user).Where("email = ?", input.Email).First(&user).Error
	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&user)
	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &user, <-errorCode
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password

	addNewUser := db.Debug().Create(&user)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &user, <-errorCode
}
