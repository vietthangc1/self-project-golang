package cache

import (
	"context"

	"github.com/thangpham4/self-project/pkg/kvredis"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
	"github.com/thangpham4/self-project/repo/mysql"
)

var _ repo.MockRepo = &MockCache{}

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
	_, _ = m.kvRedis.Get(ctx, "mock_key")
	return nil
}
