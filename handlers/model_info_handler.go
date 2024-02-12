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

type ModelInfoHandler struct {
	modelInfoService *services.ModelInfoService
	logger           logger.Logger
}

func NewModelInfoHandler(
	modelInfoService *services.ModelInfoService,
) *ModelInfoHandler {
	return &ModelInfoHandler{
		modelInfoService: modelInfoService,
		logger:           logger.Factory("ModelInfoHandler"),
	}
}

func (s *ModelInfoHandler) GetByID(
	ctx *gin.Context,
) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		const errString = "not found id in url params"
		s.logger.Error(commonx.ErrorMessages(commonx.ErrNotFoundParams, errString), errString)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		s.logger.Error(err, "error in query id from url", "id", id)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	model, err := s.modelInfoService.GetByID(ctx, uint(idInt))
	if err != nil {
		s.logger.Error(err, "error in getting model", "id", id)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, model)
}

func (s *ModelInfoHandler) GetByCode(
	ctx *gin.Context,
) {
	code, ok := ctx.Params.Get("code")
	if !ok {
		const errString = "not found code in url params"
		s.logger.Error(commonx.ErrorMessages(commonx.ErrNotFoundParams, errString), errString)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}
	model, err := s.modelInfoService.GetByCode(ctx, code)
	if err != nil {
		s.logger.Error(err, "error in getting model", "code", code)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, model)
}

func (s *ModelInfoHandler) Create(
	ctx *gin.Context,
) {
	var model entities.ModelInfo
	err := json.NewDecoder(ctx.Request.Body).Decode(&model)
	if err != nil {
		s.logger.Error(err, "error in parse json", "struct", ctx.Request.Body)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	modelPointer, err := s.modelInfoService.Create(ctx, &model)
	if err != nil {
		s.logger.Error(err, "error in create model", "model", model)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, modelPointer)
}
