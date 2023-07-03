package repo

import "context"

type MockMongoDBRepo interface {
	Get(ctx context.Context) error
}
