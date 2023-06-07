package handlerCreateProduct

import (
	"net/http"

	createProduct "github.com/donaderoyan/simple-go-api/controllers/product-controllers/create"
	util "github.com/donaderoyan/simple-go-api/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service createProduct.Service
}

func NewHandlerCreateProduct(service createProduct.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateProductHandler(ctx *gin.Context) {
	var input createProduct.InputCreateProduct

	if err := ctx.BindJSON(&input); err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}

	errResponse, err := util.Validator(input)
	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
	}

	_, errCreateProduct := h.service.CreateProductService(&input)

	switch errCreateProduct {
	case "CREATE_PRODUCT_CONFLICT_409":
		util.APIResponse(ctx, "Create product failed, user not exist", http.StatusConflict, http.MethodPost, nil)
		return
	case "CREATE_PRODUCT_FAILED_403":
		util.APIResponse(ctx, "Create product failed", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		util.APIResponse(ctx, "Create product success", http.StatusCreated, http.MethodPost, nil)
	}
}
