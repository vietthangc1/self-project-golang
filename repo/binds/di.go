package binds

import (
	"github.com/google/wire"
	"github.com/thangpham4/self-project/repo"
	"github.com/thangpham4/self-project/repo/cache"
	"github.com/thangpham4/self-project/repo/mongodb"
	"github.com/thangpham4/self-project/repo/mysql"
	"github.com/thangpham4/self-project/repo/sheet"
)

var Set = wire.NewSet(
	mysql.Set,
	cache.Set,
	mongodb.Set,
	sheet.Set,

	wire.Bind(new(repo.MockRepo), new(*cache.MockCache)),
	wire.Bind(new(repo.UserAdminRepo), new(*mysql.UserAdminMysql)),
	wire.Bind(new(repo.ProductInfoRepo), new(*cache.ProductInfoCache)),
	wire.Bind(new(repo.MockMongoDBRepo), new(*mongodb.MockMongoDB)),
	wire.Bind(new(repo.ReadModelDataRepo), new(*cache.ReadModelDataCache)),
	wire.Bind(new(repo.ModelInfoRepo), new(*mysql.ModelInfoMysql)),
)
