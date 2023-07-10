package services

import (
	"context"
	"sort"

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

func (s *BlockDataService) GetBlockProducts(
	ctx context.Context,
	blockCode, customerID string,
	pageSize, beginCursor int32,
) (*entities.BlockData, error) {
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

	productRank := modelDataCustomer.ProductRank
	sort.Slice(productRank, func(i, j int) bool {
		return productRank[i].Rank < productRank[j].Rank
	})

	productIDs := []uint{}
	for _, productID := range productRank {
		productIDs = append(productIDs, uint(productID.ProductID))
	}

	var nextCursor int32 = 0
	endCursor := beginCursor + pageSize
	if int(endCursor) < len(productIDs) {
		nextCursor = endCursor
	}

	productIDs = productIDs[beginCursor:endCursor]

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

	config := &entities.BlockDataConfig{
		BeginCursor: beginCursor,
		PageSize:    pageSize,
		BlockCode:   blockCode,
	}

	if nextCursor == 0 {
		return &entities.BlockData{
			BlockCode:  blockCode,
			ModelIDs:   modelIDs,
			Data:       productsInfo,
			ModelDebug: modelDebug,
			Config:     config,
		}, nil
	}

	moreLinkConfig := &entities.BlockDataMoreLinkConfig{
		BeginCursor: nextCursor,
		PageSize:    pageSize,
		BlockCode:   blockCode,
	}
	if customerID != "-" {
		moreLinkConfig.CustomerID = customerID
	}

	return &entities.BlockData{
		BlockCode:      blockCode,
		ModelIDs:       modelIDs,
		Data:           productsInfo,
		ModelDebug:     modelDebug,
		Config:         config,
		MoreLinkConfig: moreLinkConfig,
	}, nil
}
