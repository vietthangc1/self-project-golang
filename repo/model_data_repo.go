package repo

import (
	"context"

	"github.com/thangpham4/self-project/entities"
)

type ReadModelDataRepo interface {
	ReadModelData(
		ctx context.Context,
		blobName string,
	) ([]*entities.ModelDataMaster, error)
	ReadModelDataTransform(
		ctx context.Context,
		blobName string,
	) (map[string]*entities.ModelDataMaster, error)
}
