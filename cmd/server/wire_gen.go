// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/thangpham4/self-project/apps/server"
	"github.com/thangpham4/self-project/handlers"
	"github.com/thangpham4/self-project/infra"
	"github.com/thangpham4/self-project/pkg/kvredis"
	"github.com/thangpham4/self-project/pkg/sheets"
	"github.com/thangpham4/self-project/repo/cache"
	"github.com/thangpham4/self-project/repo/mongodb"
	"github.com/thangpham4/self-project/repo/mysql"
	"github.com/thangpham4/self-project/repo/sheet"
	"github.com/thangpham4/self-project/services"
)

// Injectors from wire.go:

func BuildServer(contextContext context.Context) (*gin.Engine, error) {
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
	mockMysql := mysql.NewMockMysql(db)
	mockCache := cache.NewMockCache(kvRedisImpl, mockMysql)
	mongoClient, err := infra.NewMongoDBConnection()
	if err != nil {
		return nil, err
	}
	mockMongoDB := mongodb.NewMockMongoDB(mongoClient)
	mockService := services.NewMockService(mockCache, mockMongoDB)
	mockHandler := handlers.NewMockHandler(mockService)
	userAdminMysql := mysql.NewUserAdminMysql(db)
	userAdminService := services.NewUserAdminService(userAdminMysql)
	userAdminHandler := handlers.NewUserAdminHandler(userAdminService)
	productInfoMysql := mysql.NewProductInfoMysql(db)
	productInfoCache := cache.NewProductInfoCache(kvRedisImpl, productInfoMysql)
	productInfoService := services.NewProductInfoService(productInfoCache)
	productInfoHandler := handlers.NewProductInfoHandler(productInfoService)
	service, err := infra.NewSheetService(contextContext)
	if err != nil {
		return nil, err
	}
	sheetService := sheets.NewSheetService(service)
	readModelSheet := sheet.NewReadModelSheet(sheetService)
	readModelDataCache := cache.NewReadModelDataCache(readModelSheet, kvRedisImpl)
	readModelDataService := services.NewReadModelDataService(readModelDataCache)
	modelInfoMysql := mysql.NewModelInfoMysql(db)
	modelInfoService := services.NewModelInfoService(modelInfoMysql)
	readModelDataHandler := handlers.NewReadModelDataHandler(readModelDataService, modelInfoService)
	modelInfoHandler := handlers.NewModelInfoHandler(modelInfoService)
	engine := server.NewHTTPserver(mockHandler, userAdminHandler, productInfoHandler, readModelDataHandler, modelInfoHandler)
	return engine, nil
}
