package services

import (
	"context"

	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/repo"
)

const (
	mockMessage string = "Hello, this is the mocked response"
)

type MockService struct {
	mockRepo      repo.MockRepo
	mockMongoRepo repo.MockMongoDBRepo
	logger        logger.Logger
}

func NewMockService(
	mockRepo repo.MockRepo,
	mockMongoRepo repo.MockMongoDBRepo,
) *MockService {
	return &MockService{
		mockRepo:      mockRepo,
		mockMongoRepo: mockMongoRepo,
		logger:        logger.Factory("MockService"),
	}
}

func (m *MockService) GetMockCache(ctx context.Context) (*entities.MockEntities, error) {
	m.logger.V(logger.LogDebugLevel).Info("Run Get Mock")

	err := m.mockRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &entities.MockEntities{
		Message: mockMessage,
	}, nil
}

func (m *MockService) GetMockMongo(ctx context.Context) (*entities.MockEntities, error) {
	m.logger.V(logger.LogDebugLevel).Info("Run Get Mock Mongo")

	err := m.mockMongoRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &entities.MockEntities{
		Message: mockMessage,
	}, nil
}
