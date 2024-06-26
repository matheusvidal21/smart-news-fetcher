package database

import (
	"database/sql"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	stmt, err := ur.db.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user models.User
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Create(user models.User) (*models.User, error) {
	stmt, err := ur.db.Prepare("INSERT INTO users (email, username, password) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	result, err := stmt.Exec(user.Email, user.Username, passwordHash)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:       int(id),
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (ur *UserRepository) Delete(email string) error {
	stmt, err := ur.db.Prepare("DELETE FROM users WHERE email = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(email)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) Update(user models.User) (*models.User, error) {
	stmt, err := ur.db.Prepare("UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(user.Username, user.Email, passwordHash, user.ID)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return &models.User{
		ID:       int(id),
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func (ur *UserRepository) FindById(id int) (*models.User, error) {
	stmt, err := ur.db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user models.User
	err = stmt.QueryRow(id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
