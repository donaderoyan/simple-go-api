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

	var paginate util.Paginate
	if err := ctx.ShouldBindQuery(&paginate); err != nil {
		// logrus.Info("GET Query Error", paginate)
	}

	resultsProduct, errResultsProduct := h.service.ResultsProductService(paginate)

	switch errResultsProduct {
	case "RESULTS_STUDENT_NOT_FOUND_404":
		util.APIResponse(ctx, "Products data is not exists", http.StatusConflict, http.MethodGet, nil)
	default:
		util.APIResponsePagination(ctx, "Results Product data successfully", http.StatusOK, http.MethodGet, resultsProduct, paginate)
	}
}
