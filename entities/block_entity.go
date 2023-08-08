package entities

import (
	"strconv"
	"strings"
)

type BlockInfo struct {
	ID          uint   `gorm:"autoIncrement" json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	ModelIDs    string `json:"model_ids,omitempty"`
}

type BlockInfoTransform struct {
	ID          uint    `gorm:"autoIncrement" json:"id,omitempty"`
	Code        string  `json:"code" gorm:"unique,omitempty"`
	Description string  `json:"description,omitempty"`
	ModelIDs    []int32 `json:"model_ids,omitempty"`
}

func (b *BlockInfo) Transform() *BlockInfoTransform {
	modelIDs := []int32{}
	modelIDString := b.ModelIDs
	if modelIDString == "" || modelIDString == "[]" {
		modelIDs = []int32{}
	}
	modelIDArr := strings.Split(modelIDString[1:len(modelIDString)-1], ",")
	for _, modelID := range modelIDArr {
		modelIDStr := strings.TrimSpace(modelID)
		modelID, err := strconv.ParseInt(modelIDStr, 10, 32)
		if err != nil {
			continue
		}
		modelIDs = append(modelIDs, int32(modelID))
	}
	return &BlockInfoTransform{
		ID:          b.ID,
		Code:        b.Code,
		Description: b.Description,
		ModelIDs:    modelIDs,
	}
}

type BlockData struct {
	BlockCode  string             `json:"block_code,omitempty"`
	ModelIDs   []int32            `json:"model_ids,omitempty"`
	Data       []*ProductInfo     `json:"data,omitempty"`
	ModelDebug *ModelDebug        `json:"model_debug,omitempty"`
	Config     *BlockDataConfig   `json:"config,omitempty"`
	MoreLink   *BlockDataMoreLink `json:"more_link,omitempty"`
}

type BlockDataConfig struct {
	BeginCursor int32  `json:"begin_cursor"`
	PageSize    int32  `json:"page_size,omitempty"`
	BlockCode   string `json:"block_code,omitempty"`
}

type BlockDataMoreLink struct {
	URL    string            `json:"url,omitempty"`
	Config map[string]string `json:"config,omitempty"`
}

type ModelDebug struct {
	Models []*ModelInfo `json:"model_info,omitempty"`
}
