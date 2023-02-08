package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"hexhoc/go-examples/internal/usecase"
)

func NewRouter(handler *gin.Engine, l *zerolog.Logger, t usecase.Product) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	//Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Routers
	h := handler.Group("/v1")
	{
		newProductRoutes(h, t, l)
	}
}
