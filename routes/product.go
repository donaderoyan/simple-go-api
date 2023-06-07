package route

import (
	createProduct "github.com/donaderoyan/simple-go-api/controllers/product-controllers/create"
	handlerCreateProduct "github.com/donaderoyan/simple-go-api/handlers/product-handlers/create"
	middleware "github.com/donaderoyan/simple-go-api/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitProductRoutes(db *gorm.DB, route *gin.Engine) {
	/**
	@description All Product Route
	*/
	createProductRepository := createProduct.NewRepositoryCreate(db)
	createProductService := createProduct.NewServiceCreate(createProductRepository)
	createProductHandler := handlerCreateProduct.NewHandlerCreateProduct(createProductService)

	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/product", createProductHandler.CreateProductHandler)
}
