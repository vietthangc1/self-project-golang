package cronjobs

import "github.com/google/wire"

var Set = wire.NewSet(
	NewProductInfoCache,
)
