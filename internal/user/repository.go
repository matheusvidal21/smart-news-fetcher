package user

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryInterface interface {
	FindByEmail(email string) (*User, error)
	Create(user User) (*User, error)
	Delete(email string) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) FindByEmail(email string) (*User, error) {
	stmt, err := ur.db.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user User
	err = stmt.QueryRow(email).Scan(&user.ID, &user.Email, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Create(user User) (*User, error) {
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

	return &User{
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
