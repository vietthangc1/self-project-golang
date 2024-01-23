package server

import (
	"github.com/gin-gonic/gin"

	"github.com/thangpham4/self-project/handlers"
	"github.com/thangpham4/self-project/middlewares"
)

func NewHTTPserver(
	mockHandler *handlers.MockHandler,
	userHandler *handlers.UserAdminHandler,
	modelHandler *handlers.ReadModelDataHandler,
	modelInfoHandler *handlers.ModelInfoHandler,
	blockInfoHanfler *handlers.BlockInfoHandler,
	blockDataHandler *handlers.BlockDataHandler,
	orderInfoHandler *handlers.OrderInfoHandler,
) *gin.Engine {
	s := gin.Default()

	s.Use(middlewares.MiddlewareUserMetaData())
	s.Use(middlewares.MiddlewareUserAdmin())

	s.GET("/mock/cache", mockHandler.GetCache)
	s.POST("/mock/kafka", mockHandler.SendMessage)
	s.GET("/mock/kafka", mockHandler.ReceiveMessage)
	s.GET("/mock/blob", mockHandler.ListBlob)

	s.POST("/user/create", userHandler.Create)
	s.GET("/user/get/:id", userHandler.Get)
	s.POST("/login", userHandler.Login)

	s.POST("/model/sheet", modelHandler.ReadModelData)

	s.POST("/model/create", modelInfoHandler.Create)
	s.GET("/model/id/:id", modelInfoHandler.GetByID)
	s.GET("/model/code/:code", modelInfoHandler.GetByCode)
	s.GET("/model/score", modelHandler.ProductScoreModelForCustomer)

	s.POST("/block/create", blockInfoHanfler.Create)
	s.GET("/block/id/:id", blockInfoHanfler.GetByID)
	s.GET("/block/code/:code", blockInfoHanfler.GetByCode)

	s.GET("/data", blockDataHandler.GetData)

	s.POST("/order/create", orderInfoHandler.Create)
	s.GET("/order/id/:id", orderInfoHandler.GetByID)
	return s
}
