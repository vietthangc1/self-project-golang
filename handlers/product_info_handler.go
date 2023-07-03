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

type ProductInfoHandler struct {
	productService *services.ProductInfoService
	logger         logger.Logger
}

func NewProductInfoHandler(
	productService *services.ProductInfoService,
) *ProductInfoHandler {
	return &ProductInfoHandler{
		productService: productService,
		logger:         logger.Factory("ProductInfoHandler"),
	}
}

func (u *ProductInfoHandler) Create(ctx *gin.Context) {
	var product entities.ProductInfo
	err := json.NewDecoder(ctx.Request.Body).Decode(&product)
	if err != nil {
		u.logger.Error(err, "error in parse json", "struct", ctx.Request.Body)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productPointer, err := u.productService.Create(ctx, &product)
	if err != nil {
		u.logger.Error(err, "error in create product", "product", product)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, productPointer)
}

func (u *ProductInfoHandler) Get(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		const errString = "not found id in url params"
		u.logger.Error(commonx.ErrorMessages(commonx.ErrNotFoundParams, errString), errString)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}

	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		u.logger.Error(err, "error in query id from url", "id", id)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	product, err := u.productService.Get(ctx, uint(idInt))
	if err != nil {
		u.logger.Error(err, "error in getting product", "id", id)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, product)
}
