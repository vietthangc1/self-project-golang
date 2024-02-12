package entities

import (
	"fmt"

	"github.com/thangpham4/self-project/pkg/commonx"
)

type ModelDataItem struct {
	ProductID int32   `json:"product_id"`
	Rank      float32 `json:"rank"`
}

type ModelDataMaster struct {
	Key         string           `json:"key"`
	ProductRank []*ModelDataItem `json:"product_rank"`
}

type ModelInfo struct {
	ID            uint         `gorm:"autoIncrement" json:"id,omitempty"`
	Code          string       `json:"code" gorm:"unique"`
	ModelSourceID uint         `json:"model_source_id,omitempty"`
	Source        *ModelSource `json:"source,omitempty" gorm:"foreignKey:ModelSourceID;references:ID"`
}

func (m *ModelInfo) Validate() error {
	modelSource := m.Source
	if modelSource == nil {
		return commonx.ErrorMessages(
			commonx.ErrInsufficientDataGet,
			fmt.Sprintf("model has no source, code: %s", m.Code),
		)
	}
	blobName := modelSource.BlobName
	if blobName == "" {
		return commonx.ErrorMessages(
			commonx.ErrInsufficientDataGet,
			fmt.Sprintf("model source has nil blob name, code: %s", m.Code),
		)
	}
	return nil
}

type ModelSource struct {
	ID       uint   `gorm:"autoIncrement" json:"id"`
	BlobName string `json:"blob_name,omitempty"`
}
