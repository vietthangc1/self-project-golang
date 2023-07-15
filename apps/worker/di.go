package worker

import (
	"github.com/google/wire"
	"github.com/thangpham4/self-project/apps/worker/cronjobs"
	"github.com/thangpham4/self-project/handlers"
	"github.com/thangpham4/self-project/infra"
	"github.com/thangpham4/self-project/repo/binds"
	"github.com/thangpham4/self-project/services"
)

var ConsolidatedSet = wire.NewSet(
	infra.Set,

	binds.Set,
	services.Set,
	handlers.Set,

	Set,
)

var Set = wire.NewSet(
	cronjobs.Set,
	NewWorker,
)
