package dto

import "time"

type ContainerStatusResponseDB struct {
	IP          string    `json:"ip"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	LastSuccess time.Time `json:"lastSuccess"`
}
