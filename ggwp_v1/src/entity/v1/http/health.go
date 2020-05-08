package http

import "time"

// HealthResponse Types
type HealthResponse struct {
	Status    string     `json:"status"`
	Mysql     string     `json:"mysql"`
	Redis     string     `json:"redis"`
	CreatedAt *time.Time `json:"check_at"`
}
