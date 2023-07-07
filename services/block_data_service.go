package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type BlockDataService struct {
	blockInfoRepo   repo.BlockInfoRepo
	modelRepo       repo.ReadModelDataRepo
	modelInfoRepo   repo.ModelInfoRepo
	productInfoRepo repo.ProductInfoRepo
	logger          logger.Logger
}

func NewBlockDataService(
	blockInfoRepo repo.BlockInfoRepo,
	modelRepo repo.ReadModelDataRepo,
	modelInfoRepo repo.ModelInfoRepo,
	productInfoRepo repo.ProductInfoRepo,
) *BlockDataService {
	return &BlockDataService{
		blockInfoRepo:   blockInfoRepo,
		modelRepo:       modelRepo,
		modelInfoRepo:   modelInfoRepo,
		productInfoRepo: productInfoRepo,
		logger:          logger.Factory("BlockDataService"),
	}
}

func (s *BlockDataService) GetBlockProducts(ctx context.Context, blockCode, customerID string) (*entities.BlockData, error) {
	blockInfo, err := s.blockInfoRepo.GetByCode(ctx, blockCode)
	if err != nil {
		s.logger.Error(err, "not found block", "code", blockCode)
		return nil, err
	}

	modelIDs := blockInfo.ModelIDs
	if len(modelIDs) == 0 {
		s.logger.Error(commonx.ErrInsufficientDataGet, "block has no model", "code", blockCode)
		return nil, commonx.ErrInsufficientDataGet
	}

	modelInfo, err := s.modelInfoRepo.GetByID(ctx, uint(modelIDs[0]))
	if err != nil {
		s.logger.Error(err, "not found model", "id", modelIDs[0])
		return nil, err
	}
	err = modelInfo.Validate()
	if err != nil {
		s.logger.Error(err, "model validate fail")
	}

	sheetID, sheetName := modelInfo.Source.SheetID, modelInfo.Source.SheetName
	modelData, err := s.modelRepo.ReadModelDataTransform(ctx, sheetID, sheetName)
	if err != nil {
		s.logger.Error(err, "get model data error", "sheet_name", sheetName, "sheet_id", sheetID)
		return nil, err
	}

	var modelDataCustomer *entities.ModelDataMaster
	modelDataCustomer, ok := modelData[customerID]
	if !ok {
		s.logger.Info(
			"customer id not in pool model, return default data",
			"customer_id", customerID,
			"model_code", modelInfo.Code,
		)
		modelDataCustomer = modelData["-"]
	}

	productIDs := []uint{}
	for _, productID := range modelDataCustomer.ProductRank {
		productIDs = append(productIDs, uint(productID.ProductID))
	}
	productsInfo, err := s.productInfoRepo.GetMany(ctx, productIDs)
	if err != nil {
		s.logger.Error(err, "cannot get products info")
		return nil, err
	}

	modelDebug := &entities.ModelDebug{
		Models: []*entities.ModelInfo{
			modelInfo,
		},
	}

	return &entities.BlockData{
		BlockCode:  blockCode,
		ModelIDs:   modelIDs,
		Data:       productsInfo,
		ModelDebug: modelDebug,
	}, nil
}
