package mysql

import (
	"context"
	"fmt"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"gorm.io/gorm"
)

type ModelInfoMysql struct {
	db *gorm.DB
}

func NewModelInfoMysql(
	db *gorm.DB,
) *ModelInfoMysql {
	return &ModelInfoMysql{
		db: db,
	}
}

func (m *ModelInfoMysql) GetByID(
	ctx context.Context,
	id uint,
) (*entities.ModelInfo, error) {
	var model = &entities.ModelInfo{
		ID: id,
	}
	err := m.db.WithContext(ctx).Preload("Source").First(model).Error
	if err != nil {
		return nil, commonx.ErrorMessages(err, fmt.Sprintf("cannot find model, id: %d", id))
	}
	return model, nil
}

func (m *ModelInfoMysql) GetByCode(
	ctx context.Context,
	code string,
) (*entities.ModelInfo, error) {
	var model = &entities.ModelInfo{
		Code: code,
	}
	err := m.db.WithContext(ctx).Preload("Source").First(model).Error
	if err != nil {
		return nil, commonx.ErrorMessages(err, fmt.Sprintf("cannot find model, code: %s", code))
	}
	return model, nil
}

func (m *ModelInfoMysql) Create(ctx context.Context, model *entities.ModelInfo) (*entities.ModelInfo, error) {
	err := m.db.WithContext(ctx).Create(model).Error
	if err != nil {
		return nil, commonx.ErrorMessages(err, fmt.Sprintf("cannot create model: %v", model))
	}
	return model, nil
}
