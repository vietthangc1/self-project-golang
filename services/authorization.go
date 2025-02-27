package services

import (
	"context"
	"strconv"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
)

type AuthorizationService struct {
	userAdminService *UserAdminService
	logger           logger.Logger
}

func NewAuthorizationService(
	userAdminService *UserAdminService,
) *AuthorizationService {
	return &AuthorizationService{
		userAdminService: userAdminService,
		logger:           logger.Factory("AuthorizationService"),
	}
}

func (s *AuthorizationService) VerifyMetaData(
	ctx context.Context,
) (*entities.UserAdmin, error) {
	userAdminContex := ctx.Value(commonx.UserAdminCtxKey)
	if userAdminContex == nil {
		s.logger.Error(commonx.ErrNotAuthenticated, "cannot find user info")
		return nil, commonx.ErrNotAuthenticated
	}
	userAdminInfo := userAdminContex.(*entities.UserAdminData)
	var user entities.UserAdmin
	user, err := s.userAdminService.GetByEmail(ctx, userAdminInfo.Email)
	if err != nil {
		s.logger.Error(err, "cannot find this email", "email", userAdminInfo.Email)
		return nil, err
	}
	return &user, nil
}

func (s *AuthorizationService) GetCustomerInfo(
	ctx context.Context,
) (int32, error) {
	customerContex := ctx.Value(commonx.UserMetadataCtxKey)
	if customerContex == nil {
		s.logger.Error(commonx.ErrUnknown, "cannot find customer id")
		return 0, commonx.ErrUnknown
	}
	customerInfo := customerContex.(*entities.UserMetadata)
	customerID := customerInfo.CustomerID
	if customerID != "" {
		customerIDInt, err := strconv.ParseInt(customerID, 10, 64)
		if err != nil {
			s.logger.Error(err, "error in parsing int", "customer_id", customerIDInt)
			return 0, err
		}
		return int32(customerIDInt), nil
	}
	return 0, commonx.ErrUnknown
}
