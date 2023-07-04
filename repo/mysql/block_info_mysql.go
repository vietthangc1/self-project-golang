package mysql

import (
	"context"
	"fmt"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/repo"
	"gorm.io/gorm"
)

var _ repo.BlockInfoRepo = &BlockInfoMysql{}

type BlockInfoMysql struct {
	db *gorm.DB
}

func NewBlockInfoMysql(
	db *gorm.DB,
) *BlockInfoMysql {
	return &BlockInfoMysql{
		db: db,
	}
}

func (b *BlockInfoMysql) Get(ctx context.Context, id uint) (*entities.BlockInfo, error) {
	var block = &entities.BlockInfo{
		ID: id,
	}
	err := b.db.WithContext(ctx).First(block).Error
	if err != nil {
		return nil, commonx.ErrorMessages(err, fmt.Sprintf("cannot find block, id: %d", id))
	}
	return block, nil
}

func (b *BlockInfoMysql) GetByCode(ctx context.Context, code string) (*entities.BlockInfo, error) {
	var block = &entities.BlockInfo{
		Code: code,
	}
	err := b.db.WithContext(ctx).First(block).Error
	if err != nil {
		return nil, commonx.ErrorMessages(err, fmt.Sprintf("cannot find block, code: %s", code))
	}
	return block, nil
}

func (b *BlockInfoMysql) Create(ctx context.Context, block *entities.BlockInfo) (*entities.BlockInfo, error) {
	err := b.db.WithContext(ctx).Create(block).Error
	if err != nil {
		return nil, commonx.ErrorMessages(err, "cannot create block")
	}
	return block, nil
}
