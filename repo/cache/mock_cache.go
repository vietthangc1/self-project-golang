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
	kvredis   kvredis.KVRedis
	mockMysql *mysql.MockMysql
	logger    logger.Logger
}

func NewMockCache(
	kvredis kvredis.KVRedis,
	mockMysql *mysql.MockMysql,
) *MockCache {
	return &MockCache{
		kvredis:   kvredis,
		mockMysql: mockMysql,
		logger:    logger.Factory("MockCache"),
	}
}

func (m *MockCache) Get(ctx context.Context) error {
	_, _ = m.kvredis.Get(ctx, "mock_key")
	return nil
}
