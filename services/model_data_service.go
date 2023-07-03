package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type ReadModelDataService struct {
	modelRepo repo.ReadModelDataRepo
	logger    logger.Logger
}

func NewReadModelDataService(
	modelRepo repo.ReadModelDataRepo,
) *ReadModelDataService {
	return &ReadModelDataService{
		modelRepo: modelRepo,
		logger:    logger.Factory("ReadModelDataService"),
	}
}

func (s *ReadModelDataService) ReadModelData(
	ctx context.Context,
	sheetID, sheetName string,
) ([]*entities.ModelDataMaster, error) {
	return s.modelRepo.ReadModelData(ctx, sheetID, sheetName)
}
