package mysql

import "github.com/google/wire"

var Set = wire.NewSet(
	NewMockMysql,
	NewUserAdminMysql,
	NewProductInfoMysql,
	NewModelInfoMysql,
	NewBlockInfoMysql,
)
