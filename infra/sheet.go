package infra

import (
	"context"

	"github.com/thangpham4/self-project/pkg/logger"
	"google.golang.org/api/sheets/v4"
)

func NewSheetService(ctx context.Context) (*sheets.Service, error) {
	l := logger.Factory("Setup Sheet service")
	service, err := sheets.NewService(ctx)
	if err != nil {
		l.V(logger.LogErrorLevel).Error(err, "failed to set up sheet service")
		return nil, err
	}
	l.V(logger.LogInfoLevel).Info("successfully set up sheet service")
	return service, nil
}
