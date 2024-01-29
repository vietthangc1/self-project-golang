package infra

import (
	"github.com/google/wire"
	"github.com/thangpham4/self-project/pkg/apix"
	"github.com/thangpham4/self-project/pkg/blobx"
	"github.com/thangpham4/self-project/pkg/kvredis"
)

var Set = wire.NewSet(
	NewMySQLConnection,

	NewRedisConfig,
	NewRedisClient,

	NewBlobConnection,

	blobx.NewBlobService,
	wire.Bind(new(blobx.BlobService), new(*blobx.BlobServiceImpl)),

	kvredis.NewKVRedis,
	wire.Bind(new(kvredis.KVRedis), new(*kvredis.KVRedisImpl)),

	apix.NewAPICaller,
	wire.Bind(new(apix.APICaller), new(*apix.APICallerImpl)),
)
