package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	tokens "github.com/thangpham4/self-project/pkg/token"
)

const (
	userAdminHeader    = "X-User-Token"
	mockUserAdminEmail = "pvthang1700@gmail.com"
)

func MiddlewareUserAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userAdminData := ExtractUserAdminFromRequest(ctx.Request)
		ctx.Set(commonx.UserAdminCtxKey, userAdminData)
		ctx.Next()
	}
}

func ExtractUserAdminFromRequest(r *http.Request) interface{} {
	// generate token for debug
	token := tokens.NewToken("", mockUserAdminEmail)
	newToken, err := token.GenerateToken()
	if err != nil {
		logger.Error(err, "cannot generate new token", "email", mockUserAdminEmail)
	}

	userToken := r.Header.Get(userAdminHeader)
	var userEmail string
	if userToken != "" {
		email, err := extractValueFromToken(userToken)
		if err == nil {
			userEmail = email
			logger.DEBUG().Info("user", "user email", userEmail, "new_token", newToken)
			return &entities.UserAdminData{
				Email: email,
			}
		}
		logger.Error(err, "error in extract email from token", "token", token)
	}
	logger.DEBUG().Info("user", "new_token", newToken)
	return nil
}
