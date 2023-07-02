package binds

import (
	"github.com/google/wire"
	"github.com/thangpham4/self-project/repo"
	"github.com/thangpham4/self-project/repo/cache"
	"github.com/thangpham4/self-project/repo/mysql"
)

var Set = wire.NewSet(
	mysql.Set,
	cache.Set,

	wire.Bind(new(repo.MockRepo), new(*cache.MockCache)),
	wire.Bind(new(repo.UserAdminRepo), new(*mysql.UserAdminMysql)),
	wire.Bind(new(repo.ProductInfoRepo), new(*cache.ProductInfoCache)),
)
