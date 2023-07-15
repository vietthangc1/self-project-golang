//go:build wireinject
// +build wireinject

package main

import (
	"context"

	"github.com/google/wire"
	"github.com/thangpham4/self-project/apps/worker"
)

func BuildServer(_ context.Context) (*worker.Worker, error) {
	wire.Build(worker.ConsolidatedSet)
	return &worker.Worker{}, nil
}
