package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/VasySS/service-monitoring-vk-task/backend/internal/config"
	httpController "github.com/VasySS/service-monitoring-vk-task/backend/internal/controller/http"
	"github.com/VasySS/service-monitoring-vk-task/backend/internal/repositrory/postgres"
	"github.com/VasySS/service-monitoring-vk-task/backend/internal/usecase/statuses"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(ctx context.Context) error {
	closer := NewCloser()

	pool, err := pgxpool.New(ctx, config.C.PostgresURL)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer closer.Add(pool.Close)

	if err := pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping postgres: %w", err)
	}

	slog.Info("connected to postgres")

	pg := postgres.New(pool)

	statusesUsecase := statuses.New(pg)
	router := httpController.NewRouter(statusesUsecase)

	go startHTTP(router, closer)

	<-ctx.Done()
	slog.Info("shutting down")

	closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := closer.Close(closeCtx); err != nil {
		return err
	}

	return nil
}

func startHTTP(r http.Handler, closer *Closer) {
	addr := fmt.Sprintf("0.0.0.0:%d", config.C.Port)

	srv := &http.Server{
		Addr:         addr,
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 10,
	}
	closer.AddWithCtx(srv.Shutdown)

	slog.Info("starting http server", slog.String("addr", addr))

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to start http server: %v", err)
	}
}
