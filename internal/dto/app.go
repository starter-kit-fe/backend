package dto

import "time"

type AppVersionResponse struct {
	Now         time.Time `json:"now"`
	Version     string    `json:"version"`
	Environment string    `json:"environment"`
}
