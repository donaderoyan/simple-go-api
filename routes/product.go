package route

import (
	createProduct "github.com/donaderoyan/simple-go-api/controllers/product-controllers/create"
	resultsProduct "github.com/donaderoyan/simple-go-api/controllers/product-controllers/results"
	handlerCreateProduct "github.com/donaderoyan/simple-go-api/handlers/product-handlers/create"
	handlerResultsProduct "github.com/donaderoyan/simple-go-api/handlers/product-handlers/results"
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

	resultsProductRepository := resultsProduct.NewRepositoryResults(db)
	resultsProductService := resultsProduct.NewServiceResults(resultsProductRepository)
	resultsProductHandler := handlerResultsProduct.NewHandlerResultsProduct(resultsProductService)

	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/product", createProductHandler.CreateProductHandler)
	groupRoute.GET("/products", resultsProductHandler.ResultsProductHandler)
}
