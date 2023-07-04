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
	modelHandler *handlers.ReadModelDataHandler,
	modelInfoHandler *handlers.ModelInfoHandler,
) *gin.Engine {
	s := gin.Default()

	s.Use(middlewares.MiddlewareUserMetaData())
	s.Use(middlewares.MiddlewareUserAdmin())

	s.GET("/mock/cache", mockHandler.GetCache)
	s.GET("/mock/mongo", mockHandler.GetMockMongo)

	s.POST("/user/create", userHandler.Create)
	s.GET("/user/get/:id", userHandler.Get)
	s.POST("/login", userHandler.Login)

	s.POST("/product/create", productHandler.Create)
	s.GET("/product/get/:id", productHandler.Get)

	s.POST("/model/sheet", modelHandler.ReadModelData)

	s.POST("/model/create", modelInfoHandler.Create)
	s.GET("/model/id/:id", modelInfoHandler.GetByID)
	s.GET("/model/code/:code", modelInfoHandler.GetByCode)
	s.GET("/model/score", modelHandler.ProductScoreModelForCustomer)
	return s
}
