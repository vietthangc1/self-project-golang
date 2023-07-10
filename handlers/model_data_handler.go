package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type ReadModelDataHandler struct {
	modelService     *services.ReadModelDataService
	modelInfoService *services.ModelInfoService
	logger           logger.Logger
}

func NewReadModelDataHandler(
	modelService *services.ReadModelDataService,
	modelInfoService *services.ModelInfoService,
) *ReadModelDataHandler {
	return &ReadModelDataHandler{
		modelService:     modelService,
		modelInfoService: modelInfoService,
		logger:           logger.Factory("ReadModelDataHandler"),
	}
}

func (h *ReadModelDataHandler) ReadModelData(
	ctx *gin.Context,
) {
	var request entities.ModelSource
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

func (h *ReadModelDataHandler) ProductScoreModelForCustomer(
	ctx *gin.Context,
) {
	userMetadata := ctx.Value(commonx.UserMetadataCtxKey).(*entities.UserMetadata)
	customerID := "-"
	if userMetadata.CustomerID != "" {
		customerID = userMetadata.CustomerID
	}

	params := ctx.Request.URL.Query()
	modelCode := params.Get("model")

	modelScore, modelInfo, err := h.modelService.ReadModelDataForCustomerFromCode(ctx, modelCode, customerID)
	if err != nil {
		h.logger.Error(err, "error in getting model score for customer", "model_code", modelCode, "customer_id", customerID)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"model_score": modelScore,
		"model_debug": modelInfo,
	})
}
