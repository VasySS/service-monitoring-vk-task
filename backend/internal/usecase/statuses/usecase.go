package statuses

import (
	"context"

	"github.com/VasySS/service-monitoring-vk-task/backend/internal/dto"
	"github.com/VasySS/service-monitoring-vk-task/backend/internal/entity"
)

type Repository interface {
	NewContainerStatuses(ctx context.Context, statuses []entity.ContainerStatus) error
	GetContainerStatuses(ctx context.Context) ([]dto.ContainerStatusResponseDB, error)
}

type Usecase struct {
	repo Repository
}

func New(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) NewContainerStatuses(ctx context.Context, statuses []entity.ContainerStatus) error {
	return u.repo.NewContainerStatuses(ctx, statuses)
}

func (u *Usecase) ContainerStatuses(ctx context.Context) ([]dto.ContainerStatusResponseDB, error) {
	return u.repo.GetContainerStatuses(ctx)
}
