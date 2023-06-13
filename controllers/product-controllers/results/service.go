package resultsProduct

import model "github.com/donaderoyan/simple-go-api/models"

type Service interface {
	ResultsProductService() (*[]model.Product, string)
}

type service struct {
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultsProductService() (*[]model.Product, string) {
	resultsProduct, errResultsProduct := s.repository.ResultsProductRepository()
	return resultsProduct, errResultsProduct
}
