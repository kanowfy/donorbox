package cron

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/kanowfy/donorbox/internal/service"
)

// CronJob contains dependencies for running cronjob.
type CronJob struct {
	projectService service.Project
	scheduler      *gocron.Scheduler
}

// New creates a new CronJob instance.
func New(service service.Project) *CronJob {
	return &CronJob{
		projectService: service,
		scheduler:      gocron.NewScheduler(time.UTC),
	}
}

// RunDaily schedules jobs to be run at 0:00 local time everyday.
func (c *CronJob) RunDaily() {
	slog.Info("Starting cronjob at daily interval...")
	s := gocron.NewScheduler(time.Local)
	s.Every(1).Day().At("00:00").StartImmediately().Do(c.checkRefutedMilestones)
	s.StartAsync()
}

func (c *CronJob) checkRefutedMilestones() {
	err := c.projectService.CheckUpdateRefutedMilestones(context.Background())
	if err != nil {
		slog.Error("error running cronjob", "error", err)
	}
}
