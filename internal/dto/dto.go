package dto

import "time"

type CreateArticleInput struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pub_date"`
	Author      string    `json:"author"`
	SourceID    int       `json:"source_id"`
}

type CreateArticleOutput struct {
	ID          string    `json:"id"`
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
	ID          string    `json:"id"`
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
	ID          string    `json:"id"`
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
	Name           string `json:"name"`
	Url            string `json:"url"`
	UserID         int    `json:"user_id"`
	UpdateInterval int    `json:"update_interval"`
}

type CreateSourceOutput struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Url            string    `json:"url"`
	UpdateInterval int       `json:"update_interval"`
	UserID         int       `json:"user_id"`
	SavedAt        time.Time `json:"saved_at"`
}

type UpdateSourceInput struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	UpdateInterval int    `json:"update_interval"`
}

type UpdateSourceOutput struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Url            string    `json:"url"`
	UpdateInterval int       `json:"update_interval"`
	UserID         int       `json:"user_id"`
	SavedAt        time.Time `json:"saved_at"`
}

type FindOneSourceOutput struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Url            string    `json:"url"`
	UpdateInterval int       `json:"update_interval"`
	UserID         int       `json:"user_id"`
	SavedAt        time.Time `json:"saved_at"`
}

type CreateUserInput struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserOutput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type FindUserByEmailOutput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserOutput struct {
	Token string `json:"token"`
}

type UpdateUserPasswordInput struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type FindUserByIdOutput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
