package sheet

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/pkg/sheets"
)

const (
	columnForKey = "A"
	columnForProductId = "B"
	columnForScore = "C"
)

type ReadModelSheet struct {
	service *sheets.SheetService
	logger logger.Logger
}

func NewReadModelSheet(
	service *sheets.SheetService,
) *ReadModelSheet {
	return &ReadModelSheet{
		service: service,
		logger: logger.Factory("ReadModelSheet"),
	}
}

func (r *ReadModelSheet) GetModelData(
	ctx context.Context,
	sheetID, sheetName string,
) (*entities.ModelDataMaster, error) {
	keyRaw, err := r.service.GetColumnData(
		ctx,
		sheetID,
		sheetName,
		columnForKey,
	)

	if err != nil {

	}

	return nil, nil
}
