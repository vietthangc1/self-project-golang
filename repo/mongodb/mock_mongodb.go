package mongodb

import (
	"context"

	"github.com/thangpham4/self-project/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ repo.MockMongoDBRepo = &MockMongoDB{}
)

type MockMongoDB struct {
	client *mongo.Client
}

func NewMockMongoDB(
	client *mongo.Client,
) *MockMongoDB {
	return &MockMongoDB{
		client: client,
	}
}

func (m *MockMongoDB) Get(ctx context.Context) error {
	return nil
}
