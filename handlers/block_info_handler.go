package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type BlockInfoHandler struct {
	blockInfoService *services.BlockInfoService
	logger           logger.Logger
}

func NewBlockInfoHanfler(
	blockInfoService *services.BlockInfoService,
) *BlockInfoHandler {
	return &BlockInfoHandler{
		blockInfoService: blockInfoService,
		logger:           logger.Factory("BlockInfoHandler"),
	}
}

func (h *BlockInfoHandler) GetByID(
	ctx *gin.Context,
) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		const errString = "not found id in url params"
		h.logger.Error(commonx.ErrorMessages(commonx.ErrNotFoundParams, errString), errString)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		h.logger.Error(err, "error in query id from url", "id", id)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	block, err := h.blockInfoService.GetByID(ctx, uint(idInt))
	if err != nil {
		h.logger.Error(err, "error in getting model", "id", id)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, block)
}

func (h *BlockInfoHandler) GetByCode(
	ctx *gin.Context,
) {
	code, ok := ctx.Params.Get("code")
	if !ok {
		const errString = "not found code in url params"
		h.logger.Error(commonx.ErrorMessages(commonx.ErrNotFoundParams, errString), errString)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}
	block, err := h.blockInfoService.GetByCode(ctx, code)
	if err != nil {
		h.logger.Error(err, "error in getting model", "code", code)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, block)
}

func (h *BlockInfoHandler) Create(
	ctx *gin.Context,
) {
	var block entities.BlockInfo
	err := json.NewDecoder(ctx.Request.Body).Decode(&block)
	if err != nil {
		h.logger.Error(err, "error in parse json", "struct", ctx.Request.Body)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	blockPointer, err := h.blockInfoService.Create(ctx, &block)
	if err != nil {
		h.logger.Error(err, "error in create model", "model", block)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, blockPointer)
}
