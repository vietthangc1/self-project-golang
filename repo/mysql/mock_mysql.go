package mysql

import (
	"context"

	"github.com/thangpham4/self-project/repo"
	"gorm.io/gorm"
)

var _ repo.MockRepo = &MockMysql{}

type MockMysql struct {
	db *gorm.DB
}

func NewMockMysql(
	db *gorm.DB,
) *MockMysql {
	return &MockMysql{
		db: db,
	}
}

func (m *MockMysql) Get(ctx context.Context) error {
	return nil
}
