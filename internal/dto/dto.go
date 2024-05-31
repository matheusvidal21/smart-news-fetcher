package dto

import "time"

type CreateArticleInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pub_date"`
	Author      string    `json:"author"`
	SourceID    int       `json:"source_id"`
}

type CreateArticleOutput struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pub_date"`
	Author      string    `json:"author"`
	SourceID    int       `json:"source_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateArticleInput struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pub_date"`
	Author      string    `json:"author"`
	SourceID    int       `json:"source_id"`
}

type UpdateArticleOutput struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pub_date"`
	Author      string    `json:"author"`
	SourceID    int       `json:"source_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type FindOneArticleOutput struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pub_date"`
	Author      string    `json:"author"`
	SourceID    int       `json:"source_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateSourceInput struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type CreateSourceOutput struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Url     string    `json:"url"`
	SavedAt time.Time `json:"saved_at"`
}

type UpdateSourceInput struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type UpdateSourceOutput struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Url     string    `json:"url"`
	SavedAt time.Time `json:"saved_at"`
}

type FindOneSourceOutput struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Url     string    `json:"url"`
	SavedAt time.Time `json:"saved_at"`
}
