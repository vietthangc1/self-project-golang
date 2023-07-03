package repo

import (
	"context"

	"github.com/thangpham4/self-project/entities"
)

type ReadModelDataRepo interface {
	GetModelData(
		ctx context.Context,
		keyRaw, productIDRaw, scoreRaw []string,
	) ([]*entities.ModelDataMaster, error)
}
