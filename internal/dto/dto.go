package dto

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var Validate = validator.New()

type CreateArticleInput struct {
	ID          string    `json:"id" validate:"required,max=36"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pub_date"`
	Author      string    `json:"author"`
	SourceID    int       `json:"source_id" validate:"required"`
}

type CreateArticleOutput struct {
	ID          string    `json:"id" validate:"required,max=36"`
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
	ID          string    `json:"id" validate:"required,max=36"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	PubDate     time.Time `json:"pub_date"`
	Author      string    `json:"author"`
	SourceID    int       `json:"source_id" validate:"required"`
}

type UpdateArticleOutput struct {
	ID          string    `json:"id" validate:"required,max=36"`
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
	Name           string `json:"name" validate:"required"`
	Url            string `json:"url" validate:"required,url"`
	UserID         int    `json:"user_id" validate:"required"`
	UpdateInterval int    `json:"update_interval" validate:"required,gt=0"`
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
	Name           string `json:"name" validate:"required"`
	Url            string `json:"url" validate:"required,url"`
	UpdateInterval int    `json:"update_interval" validate:"required,gt=0"`
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
	Username string `json:"username" validate:"required,max=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
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
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginUserOutput struct {
	Token string `json:"token"`
}

type UpdateUserPasswordInput struct {
	Email       string `json:"email" validate:"required,email"`
	OldPassword string `json:"old_password" validate:"required,min=8"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}

type FindUserByIdOutput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
