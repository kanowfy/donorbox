package cron

import (
	"log/slog"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/kanowfy/donorbox/internal/service"
)

type CronJob struct {
	projectService service.Project
	scheduler      *gocron.Scheduler
}

func New(service service.Project) *CronJob {
	return &CronJob{
		projectService: service,
		scheduler:      gocron.NewScheduler(time.UTC),
	}
}

func (c *CronJob) Start() {
	slog.Info("Starting cronjob...")
	s := gocron.NewScheduler(time.Local)
	//s.Every(1).Day().At("00:00").StartImmediately().Do(c.updateProjectStatus)
	s.StartAsync()
}

/*
func (c *CronJob) updateProjectStatus() {
	err := c.projectService.CheckAndUpdateFinishedProjects(context.Background())
	if err != nil {
		slog.Error(err.Error())
	}
}
*/
