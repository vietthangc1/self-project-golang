package services

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/thangpham4/self-project/config"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/pkg/queryx"
)

type BlockDataService struct {
	blockInfoService   *BlockInfoService
	modelService       *ReadModelDataService
	modelInfoService   *ModelInfoService
	productInfoService *ProductInfoService
	logger             logger.Logger
}

func NewBlockDataService(
	blockInfoService *BlockInfoService,
	modelService *ReadModelDataService,
	modelInfoService *ModelInfoService,
	productInfoService *ProductInfoService,
) *BlockDataService {
	return &BlockDataService{
		blockInfoService:   blockInfoService,
		modelService:       modelService,
		modelInfoService:   modelInfoService,
		productInfoService: productInfoService,
		logger:             logger.Factory("BlockDataService"),
	}
}

func (s *BlockDataService) GetModelsByBlockCode(
	ctx context.Context,
	blockCode string,
) ([]*entities.ModelInfo, error) {
	models := []*entities.ModelInfo{}

	blockInfo, err := s.blockInfoService.GetByCode(ctx, blockCode)
	if err != nil {
		s.logger.Error(err, "not found block", "code", blockCode)
		return nil, err
	}

	modelIDs := blockInfo.ModelIDs
	if len(modelIDs) == 0 {
		s.logger.Error(commonx.ErrInsufficientDataGet, "block has no model", "code", blockCode)
		return nil, commonx.ErrInsufficientDataGet
	}

	for _, modelID := range modelIDs {
		modelInfo, err2 := s.modelInfoService.GetByID(ctx, uint(modelID))
		if err2 != nil {
			s.logger.Error(err2, "not found model", "id", modelIDs[0])
		}
		models = append(models, modelInfo)
	}

	if len(models) == 0 {
		s.logger.Error(commonx.ErrItemNotFound, "not found models", "model_ids", modelIDs)
		return nil, err
	}
	return models, nil
}

//nolint:funlen,nestif
func (s *BlockDataService) GetBlockProducts(
	ctx context.Context,
	pageToken, blockCode, customerID string,
	pageSize, beginCursor int32,
) (*entities.BlockData, error) {
	if pageToken != "" {
		var err error
		queryMap, err := queryx.ReadMoreLink(pageToken)
		if err != nil {
			s.logger.Error(err, "error in read page token")
		} else {
			s.logger.V(logger.LogDebugLevel).Info("page token query", "query_map", queryMap)
			newBeginCursor, ok := queryMap["begin_cursor"]
			if ok {
				newBeginCursor, err := strconv.ParseInt(newBeginCursor, 10, 32)
				if err == nil {
					beginCursor = int32(newBeginCursor)
				}
			}
		}
	}

	models, err := s.GetModelsByBlockCode(ctx, blockCode)
	if err != nil {
		s.logger.Error(err, "err in get models", "block_code", blockCode)
	}
	modelInfo := models[0]
	err = modelInfo.Validate()
	if err != nil {
		s.logger.Error(err, "model validate fail")
	}

	modelIDs := []int32{}
	for _, model := range models {
		modelIDs = append(modelIDs, int32(model.ID))
	}

	modelDataCustomer, err := s.modelService.ReadModelDataForCustomerFromEntity(ctx, modelInfo, customerID)
	if err != nil {
		s.logger.Error(err, "error in read model data for customer")
		return nil, err
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

	productsInfo, err := s.productInfoService.GetMany(ctx, productIDs)
	if err != nil {
		s.logger.Error(err, "cannot get products info")
		return nil, err
	}

	modelDebug := &entities.ModelDebug{
		Models: []*entities.ModelInfo{
			modelInfo,
		},
	}

	blockDataConfig := &entities.BlockDataConfig{
		BeginCursor: beginCursor,
		PageSize:    pageSize,
		BlockCode:   blockCode,
	}

	resp := &entities.BlockData{
		BlockCode:  blockCode,
		ModelIDs:   modelIDs,
		Data:       productsInfo,
		ModelDebug: modelDebug,
		Config:     blockDataConfig,
	}

	if nextCursor == 0 {
		return resp, nil
	}

	moreLinkToken, _ := queryx.GenerateMoreLink(map[string]string{
		"begin_cursor": fmt.Sprintf("%d", nextCursor),
	})

	moreLinkMap := make(map[string]string)
	moreLinkMap["page_size"] = fmt.Sprintf("%d", pageSize)
	moreLinkMap["block_code"] = blockCode
	if customerID != "-" {
		moreLinkMap["customer_id"] = customerID
	}
	queryMapStr := queryx.BuildFromMap(moreLinkMap)
	moreLinkURL := fmt.Sprintf("%s/data?%s&%s", config.Domain, queryMapStr, moreLinkToken)

	moreLink := &entities.BlockDataMoreLink{
		Config: moreLinkMap,
		URL:    moreLinkURL,
	}
	resp.MoreLink = moreLink
	return resp, nil
}
