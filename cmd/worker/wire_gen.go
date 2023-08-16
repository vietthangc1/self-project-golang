// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/thangpham4/self-project/apps/worker"
	"github.com/thangpham4/self-project/apps/worker/cronjobs"
	"github.com/thangpham4/self-project/apps/worker/kafka"
	"github.com/thangpham4/self-project/infra"
	"github.com/thangpham4/self-project/pkg/kvredis"
	"github.com/thangpham4/self-project/repo/cache"
	"github.com/thangpham4/self-project/repo/mysql"
	"github.com/thangpham4/self-project/services"
)

// Injectors from wire.go:

func BuildServer(contextContext context.Context) (*worker.Worker, error) {
	redisConfig := infra.NewRedisConfig()
	client, err := infra.NewRedisClient(contextContext, redisConfig)
	if err != nil {
		return nil, err
	}
	kvRedisImpl := kvredis.NewKVRedis(client)
	db, err := infra.NewMySQLConnection()
	if err != nil {
		return nil, err
	}
	productInfoMysql := mysql.NewProductInfoMysql(db)
	productInfoCache := cache.NewProductInfoCache(kvRedisImpl, productInfoMysql)
	productInfoService := services.NewProductInfoService(productInfoCache)
	productInfoCronCache := cronjobs.NewProductInfoCache(contextContext, productInfoService)
	orderKafka, err := kafka.NewOrderKafka()
	if err != nil {
		return nil, err
	}
	workerWorker := worker.NewWorker(productInfoCronCache, orderKafka)
	return workerWorker, nil
}
