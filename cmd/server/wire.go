//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/thangpham4/self-project/apps/server"

	"github.com/gin-gonic/gin"
)

func BuildServer(_ context.Context) (*gin.Engine, error) {
	wire.Build(server.ConsolidatedSet)
	return &gin.Engine{}, nil
}
