package route

import (
	createCategory "github.com/donaderoyan/simple-go-api/controllers/category-controllers/create"
	handlerCreateCategory "github.com/donaderoyan/simple-go-api/handlers/category-handlers/create"
	middleware "github.com/donaderoyan/simple-go-api/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCategoryRoutes(db *gorm.DB, route *gin.Engine) {
	/**
	@description All Category Route
	*/
	createCategoryRepository := createCategory.NewRepositoryCreate(db)
	createCategoryService := createCategory.NewServiceCreate(createCategoryRepository)
	createCategoryHandler := handlerCreateCategory.NewHandlerCreateCategory(createCategoryService)

	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/category", createCategoryHandler.CreateCategoryHandler)
}
