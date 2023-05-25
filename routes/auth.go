package route

import (
	loginAuth "github.com/donaderoyan/simple-go-api/controllers/auth-controllers/login"
	registerAuth "github.com/donaderoyan/simple-go-api/controllers/auth-controllers/register"
	handlerLogin "github.com/donaderoyan/simple-go-api/handlers/auth-handlers/login"
	handlerRegister "github.com/donaderoyan/simple-go-api/handlers/auth-handlers/register"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerHandler := handlerRegister.NewHandlerRegister(registerService)

	loginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(loginRepository)
	loginHandler := handlerLogin.NewHandlerLogin(loginService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerHandler.RegisterHandler)
	groupRoute.POST("/login", loginHandler.LoginHandler)
}
