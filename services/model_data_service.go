package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

type ReadModelDataService struct {
	modelRepo     repo.ReadModelDataRepo
	modelInfoRepo repo.ModelInfoRepo
	logger        logger.Logger
}

func NewReadModelDataService(
	modelRepo repo.ReadModelDataRepo,
	modelInfoRepo repo.ModelInfoRepo,
) *ReadModelDataService {
	return &ReadModelDataService{
		modelRepo:     modelRepo,
		modelInfoRepo: modelInfoRepo,
		logger:        logger.Factory("ReadModelDataService"),
	}
}

func (s *ReadModelDataService) ReadModelData(
	ctx context.Context,
	sheetID, sheetName string,
) ([]*entities.ModelDataMaster, error) {
	return s.modelRepo.ReadModelData(ctx, sheetID, sheetName)
}

func (s *ReadModelDataService) ReadModelDataForCustomerFromEntity(
	ctx context.Context,
	modelInfo *entities.ModelInfo,
	customerID string,
) (*entities.ModelDataMaster, error) {
	sheetID, sheetName := modelInfo.Source.SheetID, modelInfo.Source.SheetName
	modelData, err := s.modelRepo.ReadModelDataTransform(ctx, sheetID, sheetName)
	if err != nil {
		s.logger.Error(err, "get model data error", "sheet_name", sheetName, "sheet_id", sheetID)
		return nil, err
	}

	var modelDataCustomer *entities.ModelDataMaster
	modelDataCustomer, ok := modelData[customerID]
	if !ok {
		s.logger.Info(
			"customer id not in pool model, return default data",
			"customer_id", customerID,
			"model_code", modelInfo.Code,
		)
		modelDataCustomer = modelData["-"]
	}
	return modelDataCustomer, nil
}

func (s *ReadModelDataService) ReadModelDataForCustomerFromCode(
	ctx context.Context,
	modelCode, customerID string,
) (*entities.ModelDataMaster, *entities.ModelInfo, error) {
	model, err := s.modelInfoRepo.GetByCode(ctx, modelCode)
	if err != nil {
		s.logger.Error(err, "error in getting model", "code", "code")
		return nil, nil, err
	}

	err = model.Validate()
	if err != nil {
		s.logger.Error(err, "validate model fail", "code", modelCode)
		return nil, model, err
	}

	sheetID, sheetName := model.Source.SheetID, model.Source.SheetName

	modelData, err := s.modelRepo.ReadModelDataTransform(ctx, sheetID, sheetName)
	if err != nil {
		s.logger.Error(err, "error in reading model data", "sheet_id", sheetID, "sheet_name", sheetName)
		return nil, model, err
	}

	var (
		customerIDDefault = "-"
	)

	var modelDataCustomer *entities.ModelDataMaster
	modelDataCustomer, ok := modelData[customerID]
	if !ok {
		s.logger.Info(
			"customer id not in pool model, return default data",
			"customer_id", customerID,
			"sheet_id", sheetID,
			"sheet_name", sheetName,
		)
		modelDataCustomer = modelData[customerIDDefault]
	}
	return modelDataCustomer, model, nil
}
