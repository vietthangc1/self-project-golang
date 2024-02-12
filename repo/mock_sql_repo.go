package repo

import "context"

type MockRepo interface {
	Get(ctx context.Context) error
}
