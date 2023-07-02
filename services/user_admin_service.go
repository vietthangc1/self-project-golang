package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type UserAdminService struct {
	userRepo repo.UserAdminRepo
	logger logger.Logger
}

func NewUserAdminService(
	userRepo repo.UserAdminRepo,
) *UserAdminService {
	return &UserAdminService{
		userRepo: userRepo,
		logger: logger.Factory("UserAdminService"),
	}
}

func (u *UserAdminService) Create(ctx context.Context, user entities.UserAdmin) (entities.UserAdmin, error) {
	return u.userRepo.Create(ctx, user)
}

func (u *UserAdminService) Get(ctx context.Context, id uint) (entities.UserAdmin, error) {
	return u.userRepo.Get(ctx, id)
}
