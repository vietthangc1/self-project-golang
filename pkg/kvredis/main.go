package kvredis

import (
	"context"
	"errors"
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
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	GetMany(ctx context.Context, keys []string) (map[string][]byte, []string, error)
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
	if errors.Is(err, redis.Nil) {
		r.logger.Info("key not found", "key", key)
		return nil, commonx.ErrKeyNotFound
	}
	return buf, err
}

//nolint:gocritic
func (r *KVRedisImpl) GetMany(ctx context.Context, keys []string) (map[string][]byte, []string, error) {
	inValidKeys := []string{}

	pipe := r.client.Pipeline()
	cmds := map[string]*redis.StringCmd{}

	for _, key := range keys {
		cmds[key] = pipe.Get(ctx, key)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		r.logger.Error(err, "redis exec error")
		return nil, nil, err
	}

	out := map[string][]byte{}
	for k, cmd := range cmds {
		buf, err := cmd.Bytes()
		if err != nil {
			inValidKeys = append(inValidKeys, k)
			continue
		}
		out[k] = buf
	}

	return out, inValidKeys, nil
}

func (r *KVRedisImpl) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}
