package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type BlockDataHandler struct {
	blockDataService *services.BlockDataService
	logger           logger.Logger
}

func NewBlockDataHandler(
	blockDataService *services.BlockDataService,
) *BlockDataHandler {
	return &BlockDataHandler{
		blockDataService: blockDataService,
		logger:           logger.Factory("BlockDataHandler"),
	}
}

func (h *BlockDataHandler) GetData(ctx *gin.Context) {
	userMetadata := ctx.Value(commonx.UserMetadataCtxKey).(*entities.UserMetadata)
	customerID := "-"
	if userMetadata.CustomerID != "" {
		customerID = userMetadata.CustomerID
	}

	params := ctx.Request.URL.Query()
	blockCode := params.Get("block_code")

	resp, err := h.blockDataService.GetBlockProducts(ctx, blockCode, customerID)
	if err != nil {
		h.logger.Error(err, "error in get block product", "block_code", blockCode, "customer_id", customerID)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, resp)
}
