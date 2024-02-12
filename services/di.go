package services

import "github.com/google/wire"

var Set = wire.NewSet(
	NewMockService,
	NewUserAdminService,
	NewProductInfoService,
	NewReadModelDataService,
	NewModelInfoService,
	NewBlockInfoService,
	NewBlockDataService,
	NewAuthorizationService,
	NewOrderInfoService,
)
