package server

import (
	"github.com/gin-gonic/gin"

	"github.com/thangpham4/self-project/handlers"
	"github.com/thangpham4/self-project/middlewares"
)

func NewHTTPserver(
	mockHandler *handlers.MockHandler,
	userHandler *handlers.UserAdminHandler,
	productHandler *handlers.ProductInfoHandler,
) *gin.Engine {
	s := gin.Default()

	s.Use(middlewares.MiddlewareUserMetaData())
	s.Use(middlewares.MiddlewareUserAdmin())

	s.GET("/mock", mockHandler.Get)

	s.POST("/user/create", userHandler.Create)
	s.GET("/user/get/:id", userHandler.Get)

	s.POST("/product/create", productHandler.Create)
	s.GET("/product/get/:id", productHandler.Get)

	return s
}
