package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type ReadModelDataService struct {
	modelRepo repo.ReadModelDataRepo
	logger    logger.Logger
}

func NewReadModelDataService(
	modelRepo repo.ReadModelDataRepo,
) *ReadModelDataService {
	return &ReadModelDataService{
		modelRepo: modelRepo,
		logger:    logger.Factory("ReadModelDataService"),
	}
}

func (s *ReadModelDataService) ReadModelData(
	ctx context.Context,
	sheetID, sheetName string,
) ([]*entities.ModelDataMaster, error) {
	return s.modelRepo.ReadModelData(ctx, sheetID, sheetName)
}

func (s *ReadModelDataService) ReadModelDataForCustomer(
	ctx context.Context,
	sheetID, sheetName, customerID string,
) (*entities.ModelDataMaster, error) {
	modelData, err := s.ReadModelData(ctx, sheetID, sheetName)
	if err != nil {
		s.logger.Error(err, "error in reading model data", "sheet_id", sheetID, "sheet_name", sheetName)
		return nil, err
	}

	var (
		modelForCustomer        *entities.ModelDataMaster
		modelForCustomerDefault *entities.ModelDataMaster
		isCustomerInModel       = false
		customerIDDefault       = "-"
	)

	for _, modelMaster := range modelData {
		if modelMaster.Key == customerID {
			modelForCustomer = modelMaster
			isCustomerInModel = true
			break
		}
		if modelMaster.Key == customerIDDefault {
			modelForCustomerDefault = modelMaster
		}
	}

	if !isCustomerInModel {
		s.logger.Info(
			"customer id not in pool model, return default data",
			"customer_id", customerID,
			"sheet_id", sheetID,
			"sheet_name", sheetName,
		)
		return modelForCustomerDefault, nil
	}
	return modelForCustomer, nil
}
