package createProduct

import (
	"strings"

	model "github.com/donaderoyan/simple-go-api/models"
)

type Service interface {
	CreateProductService(input *InputCreateProduct) (*model.Product, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateProductService(input *InputCreateProduct) (*model.Product, string) {

	var categories []model.Category
	var images []model.ProductImage

	strCategories := strings.Split(input.Categories, ",")
	for _, value := range strCategories {
		categories = append(categories, model.Category{ID: value})
	}

	product := model.Product{
		UserID:           input.UserID,
		ParentID:         input.ParentID,
		ProductImages:    images,
		Categories:       categories,
		Sku:              input.Sku,
		Name:             input.Name,
		Slug:             input.Slug,
		Price:            input.Price,
		Stock:            input.Stock,
		Weight:           input.Weight,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Status:           input.Status,
	}

	resultCreateProduct, errCreateProduct := s.repository.CreateProductRepository(&product)

	return resultCreateProduct, errCreateProduct

}
