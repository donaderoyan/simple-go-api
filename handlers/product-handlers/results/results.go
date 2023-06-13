package handlerResultsProduct

import (
	"net/http"

	resultsProduct "github.com/donaderoyan/simple-go-api/controllers/product-controllers/results"
	util "github.com/donaderoyan/simple-go-api/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service resultsProduct.Service
}

func NewHandlerResultsProduct(service resultsProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultsProductHandler(ctx *gin.Context) {
	resultsProduct, errResultsProduct := h.service.ResultsProductService()

	switch errResultsProduct {
	case "RESULTS_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Products data is not exists", http.StatusConflict, http.MethodPost, nil)
	default:
		util.APIResponse(ctx, "Results Product data successfully", http.StatusOK, http.MethodPost, resultsProduct)
	}
}
