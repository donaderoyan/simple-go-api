package registerAuth

import model "github.com/donaderoyan/simple-go-api/models"

type Service interface {
	RegisterService(input *InputRegister) (*model.User, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterService(input *InputRegister) (*model.User, string) {
	var user model.User

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password

	result, err := s.repository.RegisterRepository(&user)

	return result, err
}
