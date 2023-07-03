package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type MockHandler struct {
	mockService *services.MockService
	logger      logger.Logger
}

func NewMockHandler(
	mockService *services.MockService,
) *MockHandler {
	return &MockHandler{
		mockService: mockService,
		logger:      logger.Factory("MockHandler"),
	}
}

func (m *MockHandler) GetCache(ctx *gin.Context) {
	m.logger.V(logger.LogDebugLevel).Info("Running MockHandlerfor cache")
	mockEn, err := m.mockService.GetMockCache(ctx)
	mockEn.Path = ctx.Request.URL
	if err != nil {
		m.logger.Error(err, "unknown")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "err unknown"})
	}
	ctx.IndentedJSON(http.StatusFound, mockEn)
	if err != nil {
		m.logger.Error(err, "error parse json")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("error in parsing json, err: %w", err)})
	}
}

func (m *MockHandler) GetMockMongo(ctx *gin.Context) {
	m.logger.V(logger.LogDebugLevel).Info("Running MockHandlerfor cache")
	mockEn, err := m.mockService.GetMockMongo(ctx)
	mockEn.Path = ctx.Request.URL
	if err != nil {
		m.logger.Error(err, "unknown")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "err unknown"})
	}
	ctx.IndentedJSON(http.StatusFound, mockEn)
	if err != nil {
		m.logger.Error(err, "error parse json")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("error in parsing json, err: %w", err)})
	}
}
