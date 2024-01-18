package infra

import (
	"github.com/google/wire"
	"github.com/thangpham4/self-project/pkg/apix"
	"github.com/thangpham4/self-project/pkg/kvredis"
	"github.com/thangpham4/self-project/pkg/sheets"
)

var Set = wire.NewSet(
	NewMySQLConnection,
	NewSheetService,

	NewRedisConfig,
	NewRedisClient,

	NewBlobConnection,

	kvredis.NewKVRedis,
	wire.Bind(new(kvredis.KVRedis), new(*kvredis.KVRedisImpl)),

	apix.NewAPICaller,
	wire.Bind(new(apix.APICaller), new(*apix.APICallerImpl)),

	sheets.NewSheetService,
)
