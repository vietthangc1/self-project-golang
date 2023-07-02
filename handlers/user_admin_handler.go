package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type UserAdminHandler struct {
	userService *services.UserAdminService
	logger      logger.Logger
}

func NewUserAdminHandler(
	userService *services.UserAdminService,
) *UserAdminHandler {
	return &UserAdminHandler{
		userService: userService,
		logger:      logger.Factory("UserAdminHandler"),
	}
}

func (u *UserAdminHandler) Create(ctx *gin.Context) {
	var user entities.UserAdmin
	err := json.NewDecoder(ctx.Request.Body).Decode(&user)
	if err != nil {
		u.logger.Error(err, "error in parse json", "struct", ctx.Request.Body)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err = u.userService.Create(ctx, user)
	if err != nil {
		u.logger.Error(err, "error in create user", "user", user)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, user)
}

func (u *UserAdminHandler) Get(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		errString := "not found id in url params"
		u.logger.Error(fmt.Errorf(errString), errString)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": errString})
		return
	}

	idInt, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		u.logger.Error(err, "error in query id from url", "id", id)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := u.userService.Get(ctx, uint(idInt))
	if err != nil {
		u.logger.Error(err, "error in getting user", "id", id)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	ctx.IndentedJSON(http.StatusFound, user)
}
