package repo

import (
	"context"

	"github.com/thangpham4/self-project/entities"
)

type ModelInfoRepo interface {
	GetByID(
		ctx context.Context,
		id uint,
	) (*entities.ModelInfo, error)
	GetByCode(
		ctx context.Context,
		code string,
	) (*entities.ModelInfo, error)
	Create(ctx context.Context, model *entities.ModelInfo) (*entities.ModelInfo, error)
}
