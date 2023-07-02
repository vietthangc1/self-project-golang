package infra

import (
	"github.com/google/wire"
	"github.com/thangpham4/self-project/pkg/kvredis"
)

var Set = wire.NewSet(
	NewMySQLConnection,
	NewMongoDBConnection,
	NewSheetService,

	NewRedisConfig,
	NewRedisClient,

	kvredis.NewKVRedis,
	wire.Bind(new(kvredis.KVRedis), new(*kvredis.KVRedisImpl)),
)
