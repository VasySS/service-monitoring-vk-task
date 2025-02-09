package postgres

import (
	"context"
	"fmt"

	"github.com/VasySS/service-monitoring-vk-task/backend/internal/dto"
	"github.com/VasySS/service-monitoring-vk-task/backend/internal/entity"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func (p Postgres) NewContainerStatuses(ctx context.Context, statuses []entity.ContainerStatus) error {
	_, err := p.db.CopyFrom(
		ctx,
		pgx.Identifier{"container_status"},
		[]string{"ip", "status", "created_at"},
		pgx.CopyFromSlice(len(statuses), func(i int) ([]any, error) {
			return []any{
				statuses[i].IP,
				statuses[i].Status,
				statuses[i].CreatedAt,
			}, nil
		}),
	)
	if err != nil {
		return fmt.Errorf("failed to save container statuses to db: %w", err)
	}

	return nil
}

func (p Postgres) GetContainerStatuses(ctx context.Context) ([]dto.ContainerStatusResponseDB, error) {
	query := `
		SELECT DISTINCT ON (ip)
			ip, 
			status, 
			created_at, 		
			MAX(created_at) FILTER (WHERE status = 'up') OVER (PARTITION BY ip) AS last_success
		FROM container_status
		ORDER BY ip, created_at DESC
	`

	statuses := []dto.ContainerStatusResponseDB{}
	if err := pgxscan.Select(ctx, p.db, &statuses, query); err != nil {
		return nil, fmt.Errorf("failed to get container statuses from db: %w", err)
	}

	return statuses, nil
}
