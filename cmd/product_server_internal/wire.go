//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/thangpham4/self-project/apps/product_server_internal"

	"github.com/gin-gonic/gin"
)

func BuildServer(_ context.Context) (*gin.Engine, error) {
	wire.Build(productserver.ConsolidatedSet)
	return &gin.Engine{}, nil
}
