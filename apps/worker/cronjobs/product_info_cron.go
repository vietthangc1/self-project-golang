package cronjobs

import (
	"context"

	"github.com/thangpham4/self-project/pkg/cronx"
	"github.com/thangpham4/self-project/pkg/logger"
	"github.com/thangpham4/self-project/services"
)

type ProductInfoCronCache struct {
	cronJob *cronx.CronWorker
	logger  logger.Logger
}

func NewProductInfoCache(
	ctx context.Context,
	productService *services.ProductInfoService,
) *ProductInfoCronCache {
	cronWorker := cronx.NewCronWorker()

	cronWorker.RegisterJob(
		&cronx.CronJob{
			Task: func() {
				_, _ = productService.GetAll(ctx)
			},
			Scheduler: "@every 1m",
			FirstRun:  true,
		},
	)
	return &ProductInfoCronCache{
		cronJob: cronWorker,
		logger:  logger.Factory("ProductCronCache"),
	}
}

func (w *ProductInfoCronCache) Start(ctx context.Context) error {
	w.logger.Info("product cache cron worker is start")
	return w.cronJob.Start(ctx)
}

func (w *ProductInfoCronCache) Stop() {
	w.logger.Info("product cache cron worker is stop")
	w.cronJob.Stop()
}
