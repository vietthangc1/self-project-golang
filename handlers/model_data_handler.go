package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type ReadModelDataRequest struct {
	SheetName string `json:"sheet_name"`
	SheetID   string `json:"sheet_id"`
}

type ReadModelDataHandler struct {
	modelService *services.ReadModelDataService
	logger       logger.Logger
}

func NewReadModelDataHandler(
	modelService *services.ReadModelDataService,
) *ReadModelDataHandler {
	return &ReadModelDataHandler{
		modelService: modelService,
		logger:       logger.Factory("ReadModelDataHandler"),
	}
}

func (h *ReadModelDataHandler) ReadModelData(
	ctx *gin.Context,
) {
	var request ReadModelDataRequest
	err := json.NewDecoder(ctx.Request.Body).Decode(&request)
	if err != nil {
		h.logger.Error(err, "error in parse json", "struct", ctx.Request.Body)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	out, err := h.modelService.ReadModelData(ctx, request.SheetID, request.SheetName)
	if err != nil {
		h.logger.Error(err, "error in Read model data through service", "request", request)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, out)
}
