package binds

import (
	"github.com/google/wire"
	"github.com/thangpham4/self-project/repo"
	"github.com/thangpham4/self-project/repo/blob"
	"github.com/thangpham4/self-project/repo/cache"
	"github.com/thangpham4/self-project/repo/mysql"
	"github.com/thangpham4/self-project/repo/sheet"
)

var Set = wire.NewSet(
	mysql.Set,
	cache.Set,
	sheet.Set,
	blob.Set,

	wire.Bind(new(repo.MockRepo), new(*cache.MockCache)),
	wire.Bind(new(repo.UserAdminRepo), new(*mysql.UserAdminMysql)),
	wire.Bind(new(repo.ProductInfoRepo), new(*cache.ProductInfoCache)),
	wire.Bind(new(repo.ReadModelDataRepo), new(*cache.ReadModelDataCache)),
	wire.Bind(new(repo.ModelInfoRepo), new(*mysql.ModelInfoMysql)),
	wire.Bind(new(repo.BlockInfoRepo), new(*mysql.BlockInfoMysql)),
	wire.Bind(new(repo.OrderInfoRepo), new(*mysql.OrderMysql)),
)
