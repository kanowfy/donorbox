package service

import (
	"context"
	"fmt"
	"log/slog"
)

func (s *Service) CheckAndUpdateFinishedProjects(ctx context.Context) error {
	query := `UPDATE projects SET status = 'ended'
	WHERE end_date <= NOW() AND status = 'ongoing'`
	res, err := s.dbPool.Exec(ctx, query)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		slog.Info("No project needs updating", "CronjobUpdatedStatus", 0)
	} else {
		slog.Info(fmt.Sprintf("Successfully updated %d rows", res.RowsAffected()), "CronjobUpdatedStatus", res.RowsAffected())
	}

	return nil
}
