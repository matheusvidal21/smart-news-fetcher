package sources

import (
	"time"
)

type Source struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Url     string    `json:"url"`
	SavedAt time.Time `json:"saved_at"`
	UserID  int       `json:"user_id"`
}
