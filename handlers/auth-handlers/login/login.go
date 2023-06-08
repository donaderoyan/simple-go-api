package handlerLogin

import (
	"net/http"

	loginAuth "github.com/donaderoyan/simple-go-api/controllers/auth-controllers/login"
	util "github.com/donaderoyan/simple-go-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service loginAuth.Service
}

func NewHandlerLogin(service loginAuth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {
	var input loginAuth.InputLogin
	ctx.ShouldBindJSON(&input)

	errResponse, err := util.Validator(input)
	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
	}

	resultLogin, errLogin := h.service.LoginService(&input)

	switch errLogin {
	case "LOGIN_NOT_FOUND_404":
		util.APIResponse(ctx, "User account is not registered", http.StatusNotFound, http.MethodPost, nil)
		return
	case "LOGIN_WRONG_PASSWORD_403":
		util.APIResponse(ctx, "Username or password is wrong", http.StatusForbidden, http.MethodPost, nil)
		return
	default:
		accessTokenData := map[string]interface{}{"id": resultLogin.ID, "email": resultLogin.Email}
		accessToken, errToken := util.Sign(accessTokenData, "JWT_SECRET", 24*60*1)
		accessTokenData["accessToken"] = accessToken

		if errToken != nil {
			defer logrus.Error(errToken.Error())
			util.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, http.MethodPost, nil)
			return
		}

		util.APIResponse(ctx, "Login successfully", http.StatusOK, http.MethodPost, accessTokenData)
	}
}
