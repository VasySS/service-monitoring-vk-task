package app

import (
	"context"
	"log/slog"

	"github.com/VasySS/service-monitoring-vk-task/pinger/internal/usecase"
)

func Run(ctx context.Context) error {
	go usecase.ScrapeContainersLoop(ctx)

	<-ctx.Done()
	slog.Info("shutting down")

	return nil
}
