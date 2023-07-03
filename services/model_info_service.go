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
	return s.modelInfoRepo.GetByID(ctx, id)
}

func (s *ModelInfoService) GetByCode(
	ctx context.Context,
	code string,
) (*entities.ModelInfo, error) {
	return s.modelInfoRepo.GetByCode(ctx, code)
}

func (s *ModelInfoService) Create(
	ctx context.Context,
	model *entities.ModelInfo,
) (*entities.ModelInfo, error) {
	return s.modelInfoRepo.Create(ctx, model)
}
