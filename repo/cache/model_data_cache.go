package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/kvredis"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
	"github.com/thangpham4/self-project/repo/sheet"
)

var (
	modelDataCacheKeyPrefix                        = "model_data_cache_key_"
	modelCacheTTL                                  = 24 * time.Hour
	_                       repo.ReadModelDataRepo = &ReadModelDataCache{}
)

type ReadModelDataCache struct {
	modelSheet *sheet.ReadModelSheet
	kvRedis    kvredis.KVRedis
	logger     logger.Logger
}

func NewReadModelDataCache(
	modelSheet *sheet.ReadModelSheet,
	kvRedis kvredis.KVRedis,
) *ReadModelDataCache {
	return &ReadModelDataCache{
		modelSheet: modelSheet,
		kvRedis:    kvRedis,
		logger:     logger.Factory("ModelDataCache"),
	}
}

func (c *ReadModelDataCache) ReadModelData(
	ctx context.Context,
	sheetID, sheetName string,
) ([]*entities.ModelDataMaster, error) {
	cacheKey := fmt.Sprintf("%s%s-%s", modelDataCacheKeyPrefix, sheetID, sheetName)

	buf, err := c.kvRedis.Get(ctx, cacheKey)
	if err != nil {
		if !errors.Is(err, commonx.ErrKeyNotFound) {
			c.logger.Error(err, "error in get model data cache", "sheet_id", sheetID, "sheet_name", sheetName)
		}
		return c.GetandSet(ctx, sheetID, sheetName)
	}
	var modelData []*entities.ModelDataMaster
	err = json.Unmarshal(buf, &modelData)
	if err != nil {
		c.logger.Error(err, "unmarshaling cache product", "key", cacheKey)
		return nil, err
	}
	return modelData, nil
}

func (c *ReadModelDataCache) GetandSet(ctx context.Context, sheetID, sheetName string) ([]*entities.ModelDataMaster, error) {
	cacheKey := fmt.Sprintf("%s%s-%s", modelDataCacheKeyPrefix, sheetID, sheetName)
	metaLog := []interface{}{"sheet_id", sheetID, "sheet_name", sheetName}

	modelData, err := c.modelSheet.ReadModelData(ctx, sheetID, sheetName)
	if err != nil {
		c.logger.Error(err, "error in getting model data from sheet", metaLog)
		return nil, err
	}

	buf, err := json.Marshal(modelData)
	if err != nil {
		c.logger.Error(err, "err in buffering model data", metaLog)
		return modelData, nil
	}

	err = c.kvRedis.Set(ctx, cacheKey, buf, modelCacheTTL)
	if err != nil {
		c.logger.Error(err, "error in saving model data to cache", metaLog)
	}
	return modelData, nil
}
