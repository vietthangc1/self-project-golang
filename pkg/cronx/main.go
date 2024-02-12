package cronx

import (
	"context"
	"time"

	"github.com/robfig/cron"
	"github.com/thangpham4/self-project/pkg/logger"
)

type CronJob struct {
	Task      func()
	Scheduler string
	FirstRun  bool
}

type CronWorker struct {
	conductor *cron.Cron
	jobs      []*CronJob
	logger    logger.Logger
}

func NewCronWorker() *CronWorker {
	loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		loc = time.Local
	}
	return &CronWorker{
		conductor: cron.NewWithLocation(loc),
		logger:    logger.Factory("CronWorker"),
	}
}

func (c *CronWorker) RegisterJob(jobs ...*CronJob) {
	c.jobs = append(c.jobs, jobs...)
}

func (c *CronWorker) Start(ctx context.Context) error {
	c.logger.Info("cron worker star", "number of jobs", len(c.jobs))
	c.logger.Info("run for the first time")
	for _, job := range c.jobs {
		if job.FirstRun {
			go job.Task()
		}
	}
	for _, job := range c.jobs {
		if err := c.conductor.AddFunc(job.Scheduler, job.Task); err != nil {
			return err
		}
	}
	c.conductor.Start()
	return nil
}

func (c *CronWorker) Stop() {
	c.logger.Info("cron worker is stop")
	c.conductor.Stop()
}
