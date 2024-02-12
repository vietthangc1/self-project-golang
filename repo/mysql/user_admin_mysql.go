package mysql

import (
	"context"
	"fmt"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/repo"
	"gorm.io/gorm"
)

var (
	_ repo.UserAdminRepo = &UserAdminMysql{}
)

type UserAdminMysql struct {
	db *gorm.DB
}

func NewUserAdminMysql(
	db *gorm.DB,
) *UserAdminMysql {
	return &UserAdminMysql{
		db: db,
	}
}

func (u *UserAdminMysql) Create(ctx context.Context, user entities.UserAdmin) (entities.UserAdmin, error) {
	err := u.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return entities.UserAdmin{}, fmt.Errorf("cannot create user admin, user: %v, err: %w", user, err)
	}
	return user, nil
}

func (u *UserAdminMysql) Get(ctx context.Context, id uint) (entities.UserAdmin, error) {
	var user = entities.UserAdmin{
		ID: id,
	}
	err := u.db.WithContext(ctx).First(&user).Error
	if err != nil {
		return entities.UserAdmin{}, fmt.Errorf("cannot find user admin, user_id: %d, err: %w", id, err)
	}
	return user, nil
}

func (u *UserAdminMysql) GetByEmail(ctx context.Context, email string) (entities.UserAdmin, error) {
	var user = entities.UserAdmin{
		Email: email,
	}
	err := u.db.WithContext(ctx).First(&user).Error
	if err != nil {
		return entities.UserAdmin{}, commonx.ErrorMessages(err, fmt.Sprintf("cannot find user admin, email: %s", email))
	}
	return user, nil
}
