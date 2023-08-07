package resultsProduct

import (
	model "github.com/donaderoyan/simple-go-api/models"
	util "github.com/donaderoyan/simple-go-api/utils"
)

type Service interface {
	ResultsProductService(util.Paginate) (*[]model.Product, string)
}

type service struct {
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultsProductService(paginate util.Paginate) (*[]model.Product, string) {
	resultsProduct, errResultsProduct := s.repository.SetScopeResult(paginate).ResultsProductRepository()
	return resultsProduct, errResultsProduct
}
