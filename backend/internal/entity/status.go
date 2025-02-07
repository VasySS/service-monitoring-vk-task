package entity

import "time"

type ContainerStatus struct {
	ContainerID string    `json:"containerID"`
	IP          string    `json:"ip"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}
