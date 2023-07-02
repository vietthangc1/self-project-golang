package mysql

import (
	"context"
	"fmt"

	"github.com/thangpham4/self-project/entities"
	"gorm.io/gorm"
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
