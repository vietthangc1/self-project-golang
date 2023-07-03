package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type ModelInfoService struct {
	modelInfoRepo repo.ModelInfoRepo
	logger        logger.Logger
}

func NewModelInfoService(
	modelInfoRepo repo.ModelInfoRepo,
) *ModelInfoService {
	return &ModelInfoService{
		modelInfoRepo: modelInfoRepo,
		logger:        logger.Factory("ModelInfoSerevice"),
	}
}

func (s *ModelInfoService) GetByID(
	ctx context.Context,
	id uint,
) (*entities.ModelInfo, error) {

}
