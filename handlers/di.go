package handlers

import "github.com/google/wire"

var Set = wire.NewSet(
	NewMockHandler,
	NewUserAdminHandler,
	NewProductInfoHandler,
	NewReadModelDataHandler,
	NewModelInfoHandler,
)
