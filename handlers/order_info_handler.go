package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/apix"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type OrderInfoHandler struct {
	orderService     *services.OrderInfoService
	authorizeService *services.AuthorizationService
	apiClient        apix.APICaller
	logger           logger.Logger
}

func NewOrderInfoHandler(
	orderService *services.OrderInfoService,
	authorizeService *services.AuthorizationService,
	apiClient apix.APICaller,
) *OrderInfoHandler {
	return &OrderInfoHandler{
		orderService:     orderService,
		authorizeService: authorizeService,
		apiClient:        apiClient,
		logger:           logger.Factory("OrderInfoHandler"),
	}
}

func (h *OrderInfoHandler) Create(ctx *gin.Context) {
	var order entities.OrderInfoTransform
	err := json.NewDecoder(ctx.Request.Body).Decode(&order)
	if err != nil {
		h.logger.Error(err, "error in parse json", "struct", ctx.Request.Body)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderPointer, err := h.orderService.Create(ctx, &order)
	if err != nil {
		h.logger.Error(err, "error in create order", "order", order)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, orderPointer)
}

func (h *OrderInfoHandler) GetByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		const errString = "not found code in url params"
		h.logger.Error(commonx.ErrorMessages(commonx.ErrNotFoundParams, errString), errString)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}
	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		h.logger.Error(err, "error in params id from url", "id", id)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	order, err := h.orderService.GetByID(ctx, uint(idInt))
	if err != nil {
		h.logger.Error(err, "error in getting model", "id", id)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, order)
}

func (h *OrderInfoHandler) GetByCustomerID(ctx *gin.Context) {
	id, err := h.authorizeService.GetCustomerInfo(ctx)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	order, err := h.orderService.GetByCustomerID(ctx, id)
	if err != nil {
		h.logger.Error(err, "error in getting orders for customer", "customer_id", id)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, order)
}

func (h *OrderInfoHandler) GetByProductID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("product_id")
	if !ok {
		const errString = "not found product id in url params"
		h.logger.Error(commonx.ErrorMessages(commonx.ErrNotFoundParams, errString), errString)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		h.logger.Error(err, "error in params product id from url", "id", id)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	orders, err := h.orderService.GetByProductID(ctx, int32(idInt))
	if err != nil {
		h.logger.Error(err, "error in getting orders", "product_id", id)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusOK, orders)
}
