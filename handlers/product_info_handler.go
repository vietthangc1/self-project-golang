package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type ProductInfoHandler struct {
	productService       *services.ProductInfoService
	userAdminService     *services.UserAdminService
	authorizationService *services.AuthorizationService
	logger               logger.Logger
}

func NewProductInfoHandler(
	productService *services.ProductInfoService,
	userAdminService *services.UserAdminService,
	authorizationService *services.AuthorizationService,
) *ProductInfoHandler {
	return &ProductInfoHandler{
		productService:       productService,
		userAdminService:     userAdminService,
		authorizationService: authorizationService,
		logger:               logger.Factory("ProductInfoHandler"),
	}
}

func (u *ProductInfoHandler) Create(ctx *gin.Context) {
	user, err := u.authorizationService.VerifyMetaData(ctx)
	if err != nil {
		u.logger.Error(commonx.ErrNotAuthenticated, "not found user")
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": commonx.ErrNotAuthenticated.Error()})
		return
	}
	if user.Role == 0 || user.Role == 1 {
		u.logger.Error(commonx.ErrUnauthorized, "user does not have permission to create product", "user", user)
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": commonx.ErrUnauthorized.Error()})
		return
	}

	var product entities.ProductInfo
	err = json.NewDecoder(ctx.Request.Body).Decode(&product)
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

func (u *ProductInfoHandler) GetMany(ctx *gin.Context) {
	userAdminContex := ctx.Value(commonx.UserAdminCtxKey)
	if userAdminContex == nil {
		u.logger.Error(commonx.ErrNotAuthenticated, "cannot find user info")
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": commonx.ErrNotAuthenticated.Error()})
		return
	}
	userAdminInfo := userAdminContex.(*entities.UserAdminData)
	var user entities.UserAdmin
	user, err := u.userAdminService.GetByEmail(ctx, userAdminInfo.Email)
	if err != nil {
		u.logger.Error(err, "cannot find this email", "email", userAdminInfo.Email)
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if user.Role == 0 {
		u.logger.Error(commonx.ErrUnauthorized, "user does not have permission to read product", "user", user)
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": commonx.ErrUnauthorized.Error()})
		return
	}

	ids, ok := ctx.Params.Get("id")
	if !ok {
		const errString = "not found id in url params"
		u.logger.Error(commonx.ErrorMessages(commonx.ErrNotFoundParams, errString), errString)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}

	idArr := strings.Split(ids, ",")
	if len(idArr) == 0 {
		u.logger.Error(commonx.ErrKeyNotFound, "insufficient ids input", "ids", ids)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": commonx.ErrKeyNotFound})
		return
	}

	idsUint := []uint{}

	for _, idStr := range idArr {
		idInt, errParse := strconv.ParseUint(idStr, 10, 32)
		if errParse != nil {
			continue
		}
		idUint := uint(idInt)
		idsUint = append(idsUint, idUint)
	}
	products, err := u.productService.GetMany(ctx, idsUint)
	if err != nil {
		u.logger.Error(err, "error in getting products")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.IndentedJSON(http.StatusOK, products)
}
