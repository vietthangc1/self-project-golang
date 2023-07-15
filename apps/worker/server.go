package worker

import (
	"context"

	"github.com/thangpham4/self-project/apps/worker/cronjobs"
	"github.com/thangpham4/self-project/pkg/logger"
)

type Worker struct {
	productCron *cronjobs.ProductInfoCronCache
	logger      logger.Logger
}

func NewWorker(
	productCron *cronjobs.ProductInfoCronCache,
) *Worker {
	return &Worker{
		productCron: productCron,
		logger:      logger.Factory("Worker"),
	}
}

func (w *Worker) Start(ctx context.Context) error {
	return w.productCron.Start(ctx)
}

func (w *Worker) Stop() {
	w.productCron.Stop()
}
