package loginAuth

import (
	model "github.com/donaderoyan/simple-go-api/models"
)

type Service interface {
	LoginService(input *InputLogin) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) LoginService(input *InputLogin) (*model.User, string) {
	var user model.User

	user.Email = input.Email
	user.Password = input.Password

	resultLogin, errLogin := s.repository.LoginRepository(&user)
	return resultLogin, errLogin
}
