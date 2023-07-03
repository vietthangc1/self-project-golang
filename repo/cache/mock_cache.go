package cache

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/thangpham4/self-project/pkg/kvredis"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
	"github.com/thangpham4/self-project/repo/mysql"
)

var (
	_                   repo.MockRepo = &MockCache{}
	mockKey                           = "mock_key"
	mockMessageForCache               = "This is mock message for cache"
	mockTTL                           = 5 * time.Minute
)

type MockCache struct {
	kvRedis   kvredis.KVRedis
	mockMysql *mysql.MockMysql
	logger    logger.Logger
}

func NewMockCache(
	kvRedis kvredis.KVRedis,
	mockMysql *mysql.MockMysql,
) *MockCache {
	return &MockCache{
		kvRedis:   kvRedis,
		mockMysql: mockMysql,
		logger:    logger.Factory("MockCache"),
	}
}

func (m *MockCache) Get(ctx context.Context) error {
	_, err := m.kvRedis.Get(ctx, "mock_key")
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err := m.mockMysql.Get(ctx)
			if err != nil {
				m.logger.Error(err, "unknown error in cache")
				return err
			}
			err = m.kvRedis.Set(ctx, mockKey, []byte(mockMessageForCache), mockTTL)
			if err != nil {
				m.logger.Error(err, "error in set cache", "value", mockMessageForCache)
			}
			return nil
		}
	}
	return nil
}
