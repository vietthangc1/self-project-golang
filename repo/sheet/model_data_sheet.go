package sheet

import (
	"context"
	"strconv"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
)

// const (
// 	columnForKey       = "A"
// 	columnForProductID = "B"
// 	columnForScore     = "C"
// )

type ReadModelSheet struct {
	// service *sheets.SheetService
	logger logger.Logger
}

func NewReadModelSheet(
// service *sheets.SheetService,
) *ReadModelSheet {
	return &ReadModelSheet{
		// service: service,
		logger: logger.Factory("ReadModelSheet"),
	}
}

//nolint:gocritic
func (r *ReadModelSheet) GetModelRawData(
	ctx context.Context,
	sheetID, sheetName string,
) ([]string, []string, []string, error) {
	// keyRaw, err := r.service.GetColumnData(
	// 	ctx,
	// 	sheetID,
	// 	sheetName,
	// 	columnForKey,
	// )
	// if err != nil {
	// 	r.logger.Error(err, "error in read model sheet, getting key column")
	// 	return nil, nil, nil, err
	// }
	//
	// productIDRaw, err := r.service.GetColumnData(
	// 	ctx,
	// 	sheetID,
	// 	sheetName,
	// 	columnForProductID,
	// )
	// if err != nil {
	// 	r.logger.Error(err, "error in read model sheet, getting key column")
	// 	return nil, nil, nil, err
	// }
	//
	// scoreRaw, err := r.service.GetColumnData(
	// 	ctx,
	// 	sheetID,
	// 	sheetName,
	// 	columnForScore,
	// )
	// if err != nil {
	// 	r.logger.Error(err, "error in read model sheet, getting score column")
	// 	return nil, nil, nil, err
	// }
	//
	// if len(keyRaw) <= 1 || len(keyRaw) != len(productIDRaw) || len(keyRaw) != len(scoreRaw) {
	// 	r.logger.Error(
	// 		commonx.ErrInsufficientDataGet,
	// 		"key, productID, score do not have the same range",
	// 		"key", len(keyRaw),
	// 		"productID", len(productIDRaw),
	// 		"score", len(scoreRaw),
	// 	)
	// 	return nil, nil, nil, commonx.ErrorMessages(commonx.ErrInsufficientDataGet, "columns not same length")
	// }
	//
	// r.logger.V(logger.LogDebugLevel).Info("got data from sheet", "sheetId", sheetID, "len data", len(keyRaw))
	keyRaw := []string{}
	productIDRaw := []string{}
	scoreRaw := []string{}
	return keyRaw, productIDRaw, scoreRaw, nil
}

func (r *ReadModelSheet) ReadModelData(
	ctx context.Context,
	sheetID, sheetName string,
) ([]*entities.ModelDataMaster, error) {
	mapModel, err := r.ReadModelDataTransform(ctx, sheetID, sheetName)
	if err != nil {
		return nil, err
	}
	out := []*entities.ModelDataMaster{}
	for _, v := range mapModel {
		out = append(out, v)
	}
	return out, nil
}

func (r *ReadModelSheet) ReadModelDataTransform(
	ctx context.Context,
	sheetID, sheetName string,
) (map[string]*entities.ModelDataMaster, error) {
	keyRaw, productIDRaw, scoreRaw, err := r.GetModelRawData(ctx, sheetID, sheetName)
	if err != nil {
		return nil, err
	}
	// skip header
	keyRaw = keyRaw[1:]
	productIDRaw = productIDRaw[1:]
	scoreRaw = scoreRaw[1:]
	numRows := len(keyRaw)

	mapModel := map[string]*entities.ModelDataMaster{}
	keyInclude := []string{}
	for i := 0; i < numRows; i++ {
		key := keyRaw[i]
		productIDStr := productIDRaw[i]
		scoreStr := scoreRaw[i]

		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			r.logger.Error(commonx.ErrInsufficientDataGet, "wrong type of product ID", "productID", productID)
			continue
		}

		score, err := strconv.ParseFloat(scoreStr, 32)
		if err != nil {
			r.logger.Error(commonx.ErrInsufficientDataGet, "wrong type of score", "score", score)
			continue
		}

		if stringInSlice(key, keyInclude) {
			currentModelProductMaster := mapModel[key]
			currentModelProductMaster.ProductRank = append(currentModelProductMaster.ProductRank, &entities.ModelDataItem{
				ProductID: int32(productID),
				Rank:      float32(score),
			})
			continue
		}
		mapModel[key] = &entities.ModelDataMaster{
			Key: key,
			ProductRank: []*entities.ModelDataItem{{
				ProductID: int32(productID),
				Rank:      float32(score),
			}},
		}
		keyInclude = append(keyInclude, key)
	}
	return mapModel, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
