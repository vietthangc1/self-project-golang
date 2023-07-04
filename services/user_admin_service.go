package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/pkg/passwordx"
	"github.com/thangpham4/self-project/repo"
)

type UserAdminService struct {
	userRepo repo.UserAdminRepo
	logger   logger.Logger
}

func NewUserAdminService(
	userRepo repo.UserAdminRepo,
) *UserAdminService {
	return &UserAdminService{
		userRepo: userRepo,
		logger:   logger.Factory("UserAdminService"),
	}
}

func (u *UserAdminService) Create(ctx context.Context, user entities.UserAdmin) (entities.UserAdmin, error) {
	inputPassword := &passwordx.Password{
		Password: user.Password,
	}
	hasedPassword, err := inputPassword.HasingPassword()
	if err != nil {
		return entities.UserAdmin{}, err
	}
	user.Password = hasedPassword
	return u.userRepo.Create(ctx, user)
}

func (u *UserAdminService) Get(ctx context.Context, id uint) (entities.UserAdmin, error) {
	return u.userRepo.Get(ctx, id)
}

func (u *UserAdminService) GetByEmail(ctx context.Context, email string) (entities.UserAdmin, error) {
	return u.userRepo.GetByEmail(ctx, email)
}

func (u *UserAdminService) Login(ctx context.Context, email, password string) (entities.UserAdmin, error) {
	userInfo, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		u.logger.Error(err, "not found user", "email", email)
		return entities.UserAdmin{}, err
	}
	inputPassword := &passwordx.Password{
		Password: password,
	}
	ok := inputPassword.CheckPassword(userInfo.Password)
	if !ok {
		u.logger.Error(err, "wrong password", "email", email)
		return entities.UserAdmin{}, commonx.ErrInsufficientDataGet
	}
	return userInfo, nil
}
