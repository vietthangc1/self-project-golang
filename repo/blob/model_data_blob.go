package blob

import (
	"context"
	"strconv"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/blobx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

var _ repo.ReadModelDataRepo = &ReadModelBlob{}

type ReadModelBlob struct {
	blobService blobx.BlobService
	logger logger.Logger
}

func NewReadModelBlob(
	blobService blobx.BlobService,
) *ReadModelBlob {
	return &ReadModelBlob{
		blobService: blobService,
		logger: logger.Factory("ReadModelBlob"),
	}
}

func (b *ReadModelBlob) exportModelDataToList(
	ctx context.Context,
	blobName string,
) (map[string][]*entities.ModelDataItem, error) {
	csvContent, err := b.blobService.GetCSV(blobName)
	if err != nil {
		b.logger.Error(err, "Error in reading model content", "blob_name", blobName)
		return nil, err
	}
	outDict := map[string][]*entities.ModelDataItem{}
	if len(csvContent) < 2 {
		return outDict, nil
	}
	for index, row := range csvContent {
		if index > 1 {
			key := row[0]
			productID, err := strconv.ParseInt(row[1], 10, 32)
			if err != nil {
				b.logger.Error(err, "product id not int", "product_id", row[1])	
				continue
			}
			rank, err := strconv.ParseFloat(row[2], 32)
			if err != nil {
				b.logger.Error(err, "rank not int", "rank", row[2])	
				continue
			}

			newItem := &entities.ModelDataItem{
					ProductID: int32(productID),
					Rank: float32(rank),
				}

			items, ok := outDict[key]
			if ok {
				items = append(items, newItem)
			} else {
				items = []*entities.ModelDataItem{newItem}
			}
			outDict[key] = items
		}
	}
	return outDict, nil
}

func (b *ReadModelBlob) ReadModelData(
	ctx context.Context,
	blobName string,
) ([]*entities.ModelDataMaster, error) {
	out := []*entities.ModelDataMaster{}
	outDict, err := b.exportModelDataToList(ctx, blobName)
	if err != nil {
		return nil, err
	}
	for key, value := range outDict {
		out = append(out, &entities.ModelDataMaster{
			Key: key,
			ProductRank: value,
		})
	}
	return out, nil
}

func (b *ReadModelBlob) ReadModelDataTransform(
	ctx context.Context,
	blobName string,
) (map[string]*entities.ModelDataMaster, error) {
	out := map[string]*entities.ModelDataMaster{}
	outDict, err := b.exportModelDataToList(ctx, blobName)
	if err != nil {
		return nil, err
	}
	for key, value := range outDict {
		out[key] = &entities.ModelDataMaster{
			Key: key,
			ProductRank: value,
		}
	}
	return out, nil
}
