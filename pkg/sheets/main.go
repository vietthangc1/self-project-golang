package sheets

import (
	"context"
	"fmt"
	"strings"

	"github.com/thangpham4/self-project/pkg/commonx"
	"github.com/thangpham4/self-project/pkg/logger"
	"google.golang.org/api/sheets/v4"
)

type SheetService struct {
	service *sheets.Service
	logger  logger.Logger
}

func NewSheetService(
	service *sheets.Service,
) *SheetService {
	return &SheetService{
		service: service,
		logger:  logger.Factory("SheetService"),
	}
}

// GetColumnData receive sheet_name, and column name. returns array of row valune in string format
// for eg: sheet_name = "raw", column = "A" meanse the range you want to get is "raw!A:A"
func (s *SheetService) GetColumnData(ctx context.Context, sheetID, sheetName, column string) ([]string, error) {
	sheetRange := fmt.Sprintf("%s!%s:%s", sheetName, column, column)
	resp, err := s.service.Spreadsheets.Values.Get(sheetID, sheetRange).Do()
	if err != nil {
		return nil, err
	}

	sheet := resp.Values

	if len(sheet) == 0 || len(sheet) == 1 {
		return nil, commonx.ErrorMessages(commonx.ErrItemNotFound, fmt.Sprintf("not found sheet %s, sheet range %s", sheetID, sheetRange))
	}

	out := make([]string, 0, len(sheet))
	for _, row := range sheet {
		data, ok := row[0].(string)
		if !ok {
			continue
		}

		data = strings.TrimSpace(data)

		out = append(out, data)
	}

	return out, nil
}
