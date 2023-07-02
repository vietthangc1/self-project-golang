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
	"github.com/thangpham4/self-project/repo/mysql"
)

var (
	ProductCacheKeyPrefix                      = "product_cache_key_"
	ProductCacheTTL                            = 24 * time.Hour
	_                     repo.ProductInfoRepo = &ProductInfoCache{}
)

type ProductInfoCache struct {
	kvRedis      kvredis.KVRedis
	productMysql *mysql.ProductInfoMysql
	logger       logger.Logger
}

func NewProductInfoCache(
	kvRedis kvredis.KVRedis,
	productMysql *mysql.ProductInfoMysql,
) *ProductInfoCache {
	return &ProductInfoCache{
		kvRedis:      kvRedis,
		productMysql: productMysql,
		logger:       logger.Factory("ProductInfoCache"),
	}
}

func (u *ProductInfoCache) Get(ctx context.Context, id uint) (*entities.ProductInfo, error) {
	key := fmt.Sprintf("%s%d", ProductCacheKeyPrefix, id)
	buf, err := u.kvRedis.Get(ctx, key)
	if err != nil {
		if !errors.Is(err, commonx.ErrKeyNotFound) {
			u.logger.Error(err, "error in get product cache", "product_id", id)
		}
		return u.GetandSet(ctx, id)
	}
	var product *entities.ProductInfo
	err = json.Unmarshal(buf, product)
	if err != nil {
		u.logger.Error(err, "unmarshaling cache product", "key", key)
		return nil, err
	}
	return product, nil
}

func (u *ProductInfoCache) Create(ctx context.Context, product *entities.ProductInfo) (*entities.ProductInfo, error) {
	return u.productMysql.Create(ctx, product)
}

func (u *ProductInfoCache) GetandSet(ctx context.Context, id uint) (*entities.ProductInfo, error) {
	key := fmt.Sprintf("%s%d", ProductCacheKeyPrefix, id)
	product, err := u.productMysql.Get(ctx, id)
	if err != nil {
		u.logger.Error(err, "error in getting product from mysql", "id", id)
		return nil, err
	}

	buf, err := json.Marshal(product)
	if err != nil {
		u.logger.Error(err, "err in buffering product", "id", id)
		return product, nil
	}

	err = u.kvRedis.Set(ctx, key, buf, ProductCacheTTL)
	if err != nil {
		u.logger.Error(err, "error in saving to cache", "id", id)
	}
	return product, nil
}
