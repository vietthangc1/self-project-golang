package productserver

import (
	"github.com/gin-gonic/gin"

	"github.com/thangpham4/self-project/handlers"
	"github.com/thangpham4/self-project/middlewares"
)

func NewHTTPserver(
	productHandler *handlers.ProductInfoHandler,
) *gin.Engine {
	s := gin.Default()

	s.Use(middlewares.MiddlewareUserMetaData())
	s.Use(middlewares.MiddlewareUserAdmin())

	s.POST("/create", productHandler.Create)
	s.GET("/get/:id", productHandler.GetMany)

	return s
}
