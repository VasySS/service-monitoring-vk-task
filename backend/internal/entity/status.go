package entity

import "time"

type ContainerStatus struct {
	ContainerID string
	IP          string
	Status      string
	CreatedAt   time.Time
}
