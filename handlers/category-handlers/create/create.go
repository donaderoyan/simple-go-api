package handlerCreateCategory

import (
	"net/http"

	createCategory "github.com/donaderoyan/simple-go-api/controllers/category-controllers/create"
	util "github.com/donaderoyan/simple-go-api/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service createCategory.Service
}

func NewHandlerCreateCategory(service createCategory.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateCategoryHandler(ctx *gin.Context) {
	var input createCategory.InputCreateCategory

	if err := ctx.BindJSON(&input); err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}

	errResponse, err := util.Validator(input)
	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
	}

	_, errCreateCategory := h.service.CreateCategoryService(&input)

	switch errCreateCategory {
	case "CREATE_CATEGORY_FAILED_403":
		util.APIResponse(ctx, "Create category failed", http.StatusForbidden, http.MethodPost, input)
		return
	default:
		util.APIResponse(ctx, "Create category success", http.StatusCreated, http.MethodPost, input)
	}
}
