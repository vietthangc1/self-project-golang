package blob

import "github.com/google/wire"

var Set = wire.NewSet(
	NewReadModelBlob,
)
