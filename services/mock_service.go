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
	mockRepo repo.MockRepo
	logger logger.Logger
}

func NewMockService(
	mockRepo repo.MockRepo,
) *MockService {
	return &MockService{
		mockRepo: mockRepo,
		logger: logger.Factory("MockService"),
	}
}

func (m *MockService) GetMock(ctx context.Context) (*entities.MockEntities, error) {
	m.logger.V(logger.LogDebugLevel).Info("Run Get Mock")

	err := m.mockRepo.Get(ctx)
	if err != nil {
		return nil, err
	}
	return &entities.MockEntities{
		Message: mockMessage,
	}, nil
}
