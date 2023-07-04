package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type BlockInfoService struct {
	blockInfoRepo repo.BlockInfoRepo
	logger        logger.Logger
}

func NewBlockInfoService(
	blockInfoRepo repo.BlockInfoRepo,
) *BlockInfoService {
	return &BlockInfoService{
		blockInfoRepo: blockInfoRepo,
		logger:        logger.Factory("BlockInfoService"),
	}
}

func (s *BlockInfoService) Create(ctx context.Context, block *entities.BlockInfo) (*entities.BlockInfo, error) {
	return s.blockInfoRepo.Create(ctx, block)
}

func (s *BlockInfoService) GetByID(ctx context.Context, id uint) (*entities.BlockInfo, error) {
	return s.blockInfoRepo.Get(ctx, id)
}

func (s *BlockInfoService) GetByCode(ctx context.Context, code string) (*entities.BlockInfo, error) {
	return s.blockInfoRepo.GetByCode(ctx, code)
}
