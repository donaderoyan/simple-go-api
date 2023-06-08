package createCategory

import (
	model "github.com/donaderoyan/simple-go-api/models"
)

type Service interface {
	CreateCategoryService(input *InputCreateCategory) (*model.Category, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateCategoryService(input *InputCreateCategory) (*model.Category, string) {

	category := model.Category{
		Name: input.Name,
		Slug: input.Slug,

		ParentID: input.ParentID,
	}

	resultCreateCategory, errCreateCategory := s.repository.CreateCategoryRepository(&category)

	return resultCreateCategory, errCreateCategory

}
