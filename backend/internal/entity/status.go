package entity

import "time"

type ContainerStatus struct {
	IP        string    `json:"ip"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}
