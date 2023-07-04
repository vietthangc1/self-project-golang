package repo

import (
	"context"

	"github.com/thangpham4/self-project/entities"
)

type BlockInfoRepo interface {
	Get(ctx context.Context, id uint) (*entities.BlockInfo, error)
	GetByCode(ctx context.Context, code string) (*entities.BlockInfo, error)
	Create(ctx context.Context, block *entities.BlockInfo) (*entities.BlockInfo, error)
}
