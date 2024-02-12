package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/pkg/queryx"
	"github.com/thangpham4/self-project/services"
)

var (
	defaultPageSize    int32 = 8
	defaultBeginCursor int32 = 0
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
	pageToken := params.Get("page_token")

	pageSize := defaultPageSize
	err := queryx.ReadAndParseIntVariable(params, "page_size", &pageSize)
	if err != nil {
		h.logger.Error(err, "error in parse int query values")
	}

	beginCursor := defaultBeginCursor
	err = queryx.ReadAndParseIntVariable(params, "begin_cursor", &beginCursor)
	if err != nil {
		h.logger.Error(err, "error in parse int query values")
	}

	h.logger.V(logger.LogDebugLevel).Info(
		"getting block data",
		"block_code", blockCode,
		"customer_id", customerID,
		"page_size", pageSize,
		"begin_cursor", beginCursor,
	)

	resp, err := h.blockDataService.GetBlockProducts(ctx, pageToken, blockCode, customerID, pageSize, beginCursor)
	if err != nil {
		h.logger.Error(err, "error in get block product", "block_code", blockCode, "customer_id", customerID)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, resp)
}
