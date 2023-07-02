package kvredis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
)

type KVRedisImpl struct {
	client *redis.Client
	logger logger.Logger
}

type KVRedis interface {
	Get(ctx context.Context, key string) ([]byte, error) 
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) (error) 
}

var _ KVRedis = &KVRedisImpl{}

func NewKVRedis(
	client *redis.Client,
) *KVRedisImpl {
	return &KVRedisImpl{
		client: client,
		logger: logger.Factory("KVRedis"),
	}
}

func (r *KVRedisImpl) Get(ctx context.Context, key string) ([]byte, error) {
	buf, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		r.logger.Info("key not found", "key", key)
		return nil, commonx.ErrNotFound 
	}
	return buf, err
}

func (r *KVRedisImpl) Set(ctx context.Context, key string, value []byte, ttl time.Duration) (error) {
	return r.client.Set(ctx, key, value, ttl).Err()
}

