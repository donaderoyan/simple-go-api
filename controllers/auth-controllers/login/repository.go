package loginAuth

import (
	model "github.com/donaderoyan/simple-go-api/models"
	util "github.com/donaderoyan/simple-go-api/utils"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	LoginRepository(input *model.User) (*model.User, string)
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) LoginRepository(input *model.User) (*model.User, string) {
	var user model.User
	db := r.db.Model(&user)
	errorCode := make(chan string, 1)

	user.Email = input.Email
	user.Password = input.Password

	checkUserAccount := db.Debug().Model(&user).Where("email = ?", input.Email).First(&user).Error
	if checkUserAccount != nil {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &user, <-errorCode
	}

	comparePassword := util.ComparePassword(user.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &user, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &user, <-errorCode
}
