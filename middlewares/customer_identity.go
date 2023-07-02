package middlewares

import (
	"net/http"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/pkg/token"
	"github.com/thangpham4/self-project/pkg/utils"
)

const (
	customerToken = "X-Customer-Token"
	customerIdQuery = "customer_id"
	platformQuery = "platform"
)

func ExtractUserMetaDataFromRequest(r *http.Request) interface{} {
	query := r.URL.Query()

	var (
		customerIdFromToken string
		customerIdFromQuery string
)
	customerToken := r.Header.Get(customerToken)
	if customerToken != "" {
		customerId, err := extractCustomerIdFromToken(customerToken)
		if err == nil {
			customerIdFromToken = customerId
		}
	}

	customerIdFromQuery = query.Get(customerIdQuery)

	// customerId := 
	platform := query.Get(platformQuery)

	return &entities.UserMetadata{
		CustomerId: utils.Coalesce(customerIdFromToken, customerIdFromQuery, "").(int32),
		Platform: platform,
	}
}

func extractCustomerIdFromToken(s string) (string, error) {
	token := token.NewToken(s, "")

	customerId, err := token.ExtractTokenKey()
	if err != nil {
		logger.Error(err, "error in extract customer id from token", "token", token)
		return "", err
	}
	return customerId, err
}
