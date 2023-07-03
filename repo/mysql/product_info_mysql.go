package mysql

import (
	"context"
	"fmt"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/commonx"
	"gorm.io/gorm"
)

type ProductInfoMysql struct {
	db *gorm.DB
}

func NewProductInfoMysql(
	db *gorm.DB,
) *ProductInfoMysql {
	return &ProductInfoMysql{
		db: db,
	}
}

func (u *ProductInfoMysql) Create(ctx context.Context, product *entities.ProductInfo) (*entities.ProductInfo, error) {
	err := u.db.WithContext(ctx).Create(product).Error
	if err != nil {
		return nil, commonx.ErrorMessages(err, fmt.Sprintf("cannot create product"))
	}
	return product, nil
}

func (u *ProductInfoMysql) Get(ctx context.Context, id uint) (*entities.ProductInfo, error) {
	var product = &entities.ProductInfo{
		ID: id,
	}
	err := u.db.WithContext(ctx).First(product).Error
	if err != nil {
		return nil, commonx.ErrorMessages(err, fmt.Sprintf("cannot find product, id: %d", id))
	}
	return product, nil
}
