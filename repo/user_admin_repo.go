package repo

import (
	"context"

	"github.com/thangpham4/self-project/entities"
)

type UserAdminRepo interface {
	Create(ctx context.Context, user entities.UserAdmin) (entities.UserAdmin, error)
	Get(ctx context.Context, id uint) (entities.UserAdmin, error)
}
