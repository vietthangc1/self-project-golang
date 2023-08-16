package worker

import (
	"context"

	"github.com/thangpham4/self-project/apps/worker/cronjobs"
	"github.com/thangpham4/self-project/apps/worker/kafka"
	"github.com/thangpham4/self-project/pkg/logger"
)

type Worker struct {
	productCron *cronjobs.ProductInfoCronCache
	orderKafka *kafka.OrderKafka
	logger      logger.Logger
}

func NewWorker(
	productCron *cronjobs.ProductInfoCronCache,
	orderKafka *kafka.OrderKafka,
) *Worker {
	return &Worker{
		productCron: productCron,
		orderKafka: orderKafka,
		logger:      logger.Factory("Worker"),
	}
}

func (w *Worker) Start(ctx context.Context) error {
	done := make(chan error)
	go func(done chan<- error) {
		err := w.productCron.Start(ctx)
		done <- err
		err = w.orderKafka.Start(ctx)
		done <- err
	}(done)
	for err := range done {
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Worker) Stop() {
	w.productCron.Stop()
	w.orderKafka.Stop()
}
