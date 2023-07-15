package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
)

func ByPassMiddlewareUserAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(commonx.UserAdminCtxKey, &entities.UserAdminData{
			Email: mockUserAdminEmail,
		})
		ctx.Next()
	}
}
