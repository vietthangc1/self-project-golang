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

//nolint:funlen,gocyclo
func (s *BlockDataService) GetBlockProducts(
	ctx context.Context,
	pageToken, blockCode, customerID string,
	pageSize, beginCursor int32,
) (*entities.BlockData, error) {
	isUsePageToken := false
	queryMap := make(map[string]string)
	if pageToken != "" {
		var err error
		queryMap, err = queryx.ReadMoreLink(pageToken)
		if err != nil {
			s.logger.Error(err, "error in read page token")
		} else {
			s.logger.V(logger.LogDebugLevel).Info("page token query", "query_map", queryMap)
			isUsePageToken = true
		}
	}

	if isUsePageToken {
		newBeginCursor, ok := queryMap["begin_cursor"]
		if ok {
			newBeginCursor, err := strconv.ParseInt(newBeginCursor, 10, 32)
			if err == nil {
				beginCursor = int32(newBeginCursor)
			}
		}
	}
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

	moreLinkMap := make(map[string]string)
	moreLinkMap["page_size"] = fmt.Sprintf("%d", pageSize)
	moreLinkMap["block_code"] = blockCode

	if customerID != "-" {
		moreLinkMap["customer_id"] = customerID
	}
	queryMapStr := queryx.BuildFromMap(moreLinkMap)

	moreLinkMap["begin_cursor"] = fmt.Sprintf("%d", nextCursor)
	moreLinkToken, _ := queryx.GenerateMoreLink(moreLinkMap)

	moreLinkURL := fmt.Sprintf("%s/data?%s&%s", config.Domain, queryMapStr, moreLinkToken)
	moreLink := &entities.BlockDataMoreLink{
		Config: moreLinkMap,
		URL:    moreLinkURL,
	}
	resp.MoreLink = moreLink
	return resp, nil
}
