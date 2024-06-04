package models

import (
	"time"
)

type Source struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Url            string    `json:"url"`
	SavedAt        time.Time `json:"saved_at"`
	UserID         int       `json:"user_id"`
	UpdateInterval int       `json:"update_interval"`
	Subscriber     bool      `json:"subscriber"`
}
