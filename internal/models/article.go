package models

import "time"

type Article struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pub_date"`
	Author      string    `json:"author"`
	SourceID    int       `json:"source_id"`
}
