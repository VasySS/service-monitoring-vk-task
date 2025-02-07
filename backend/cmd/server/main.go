package main

import (
	"context"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/VasySS/service-monitoring-vk-task/backend/internal/app"
	"github.com/VasySS/service-monitoring-vk-task/backend/internal/config"
)

func main() {
	config.MustInit()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx); err != nil {
		slog.Error("failed to run app", slog.Any("error", err))
	}
}
