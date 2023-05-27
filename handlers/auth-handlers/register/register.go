package handlerRegister

import (
	"net/http"

	registerAuth "github.com/donaderoyan/simple-go-api/controllers/auth-controllers/register"
	util "github.com/donaderoyan/simple-go-api/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service registerAuth.Service
}

func NewHandlerRegister(service registerAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) RegisterHandler(ctx *gin.Context) {
	var input registerAuth.InputRegister
	// ctx.ShouldBindJSON(&input)

	if err := ctx.BindJSON(&input); err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err.Error())
		return
	}

	errResponse, err := util.Validator(input)
	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
	}

	_, errRegister := h.service.RegisterService(&input)

	switch errRegister {
	case "REGISTER_CONFLICT_409":
		util.APIResponse(ctx, "Email already exist", http.StatusConflict, http.MethodPost, nil)
		return
	case "REGISTER_FAILED_403":
		util.APIResponse(ctx, "Register new account failed", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		// accessTokenData := map[string]interface{}{"id": resultRegister.ID, "email": resultRegister.Email}
		// accessToken, errToken := util.Sign(accessTokenData, util.GodotEnv("JWT_SECRET"), 60)

		// if errToken != nil {
		// 	defer logrus.Error(errToken.Error())
		// 	util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
		// 	return
		// }

		// _, errSendMail := util.SendGridMail(resultRegister.FirstName, resultRegister.Email, "Activation Account", "template_register", accessToken)

		// if errSendMail != nil {
		// 	defer logrus.Error(errSendMail.Error())
		// 	util.APIResponse(ctx, "Sending email activation failed", http.StatusBadRequest, http.MethodPost, nil)
		// 	return
		// }

		util.APIResponse(ctx, "Register new account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
