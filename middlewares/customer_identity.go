package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	tokens "github.com/thangpham4/self-project/pkg/token"
	"github.com/thangpham4/self-project/pkg/utils"
)

const (
	customerToken   = "X-Customer-Token"
	customerIDQuery = "customer_id"
	platformQuery   = "platform"
)

func MiddlewareUserMetaData() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userMetadata := ExtractUserMetaDataFromRequest(ctx.Request)
		ctx.Set(commonx.UserMetadataCtxKey, userMetadata)
		ctx.Next()
	}
}

func ExtractUserMetaDataFromRequest(r *http.Request) interface{} {
	query := r.URL.Query()
	var (
		customerIDFromToken string
		customerIDFromQuery string
	)
	customerToken := r.Header.Get(customerToken)
	if customerToken != "" {
		customerID, err := extractcustomerIDFromToken(customerToken)
		if err == nil {
			customerIDFromToken = customerID
		}
	}
	customerIDFromQuery = query.Get(customerIDQuery)
	customerID := utils.Coalesce(customerIDFromToken, customerIDFromQuery, "").(string)

	// generate token for debug
	token := tokens.NewToken("", customerID)
	newToken, err := token.GenerateToken()
	if err != nil {
		logger.Error(err, "cannot generate new token", "customer_id", customerID)
	}

	platform := query.Get(platformQuery)

	logger.DEBUG().Info("user metadata", "customerID", customerID, "platform", platform, "new_token", newToken)

	return &entities.UserMetadata{
		CustomerID: customerID,
		Platform:   platform,
	}
}

func extractcustomerIDFromToken(s string) (string, error) {
	token := tokens.NewToken(s, "")

	customerID, err := token.ExtractTokenKey()
	if err != nil {
		logger.Error(err, "error in extract customer id from token", "token", token)
		return "", err
	}
	return customerID, err
}
