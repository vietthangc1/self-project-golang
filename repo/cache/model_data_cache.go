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
	"github.com/thangpham4/self-project/repo/blob"
)

var (
	modelDataCacheKeyPrefix                                 = "model_data_cache_key_"
	modelDataTransformCacheKeyPrefix                        = "model_data_transform_cache_key_"
	modelCacheTTL                                           = 24 * time.Hour
	_                                repo.ReadModelDataRepo = &ReadModelDataCache{}
)

type ReadModelDataCache struct {
	modelBlob *blob.ReadModelBlob
	kvRedis   kvredis.KVRedis
	logger    logger.Logger
}

func NewReadModelDataCache(
	modelBlob *blob.ReadModelBlob,
	kvRedis kvredis.KVRedis,
) *ReadModelDataCache {
	return &ReadModelDataCache{
		modelBlob: modelBlob,
		kvRedis:   kvRedis,
		logger:    logger.Factory("ModelDataCache"),
	}
}

func (c *ReadModelDataCache) ReadModelData(
	ctx context.Context,
	blobName string,
) ([]*entities.ModelDataMaster, error) {
	cacheKey := fmt.Sprintf("%s%s", modelDataCacheKeyPrefix, blobName)

	buf, err := c.kvRedis.Get(ctx, cacheKey)
	if err != nil {
		if !errors.Is(err, commonx.ErrKeyNotFound) {
			c.logger.Error(err, "error in get model data cache", "blob_name", blobName)
		}
		return c.GetandSet(ctx, blobName)
	}
	var modelData []*entities.ModelDataMaster
	err = json.Unmarshal(buf, &modelData)
	if err != nil {
		c.logger.Error(err, "unmarshaling cache product", "key", cacheKey)
		return nil, err
	}
	return modelData, nil
}

func (c *ReadModelDataCache) ReadModelDataTransform(
	ctx context.Context,
	blobName string,
) (map[string]*entities.ModelDataMaster, error) {
	transformCacheKey := fmt.Sprintf("%s%s", modelDataTransformCacheKeyPrefix, blobName)

	buf, err := c.kvRedis.Get(ctx, transformCacheKey)
	if err != nil {
		if !errors.Is(err, commonx.ErrKeyNotFound) {
			c.logger.Error(err, "error in get model data transform cache", "blob_name", blobName)
		}
		return c.GetandSetTransform(ctx, blobName)
	}
	var modelData map[string]*entities.ModelDataMaster
	err = json.Unmarshal(buf, &modelData)
	if err != nil {
		c.logger.Error(err, "unmarshaling cache product", "key", transformCacheKey)
		return nil, err
	}
	return modelData, nil
}

func (c *ReadModelDataCache) GetandSet(ctx context.Context, blobName string) ([]*entities.ModelDataMaster, error) {
	cacheKey := fmt.Sprintf("%s%s", modelDataCacheKeyPrefix, blobName)
	metaLog := []interface{}{"blob_name", blobName}

	modelData, err := c.modelBlob.ReadModelData(ctx, blobName)
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

func (c *ReadModelDataCache) GetandSetTransform(
	ctx context.Context,
	blobName string,
) (map[string]*entities.ModelDataMaster, error) {
	transformCacheKey := fmt.Sprintf("%s%s", modelDataTransformCacheKeyPrefix, blobName)
	metaLog := []interface{}{"blob_name", blobName}

	modelData, err := c.modelBlob.ReadModelDataTransform(ctx, blobName)
	if err != nil {
		c.logger.Error(err, "error in getting model data transform from sheet", metaLog)
		return nil, err
	}

	buf, err := json.Marshal(modelData)
	if err != nil {
		c.logger.Error(err, "err in buffering model data transform", metaLog)
		return modelData, nil
	}

	err = c.kvRedis.Set(ctx, transformCacheKey, buf, modelCacheTTL)
	if err != nil {
		c.logger.Error(err, "error in saving model data transform to cache", metaLog)
	}
	return modelData, nil
}
