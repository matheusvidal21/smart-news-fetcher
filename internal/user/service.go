package user

import (
	"errors"
	"github.com/matheusvidal21/smart-news-fetcher/internal/auth"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
)

type UserServiceInterface interface {
	FindByEmail(email string) (dto.FindUserByEmailOutput, error)
	Create(userDto dto.CreateUserInput) (dto.CreateUserOutput, error)
	Delete(email string) error
	Login(userDto dto.LoginUserInput) (dto.LoginUserOutput, error)
}

type UserService struct {
	userRepository UserRepositoryInterface
	AuthService    auth.JWTServiceInterface
}

func NewUserService(userRepository UserRepositoryInterface, authService auth.JWTServiceInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
		AuthService:    authService,
	}
}

func (us *UserService) FindByEmail(email string) (dto.FindUserByEmailOutput, error) {
	user, err := us.userRepository.FindByEmail(email)

	if err != nil {
		return dto.FindUserByEmailOutput{}, errors.New("failed to find user")
	}

	return dto.FindUserByEmailOutput{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (us *UserService) Create(userDto dto.CreateUserInput) (dto.CreateUserOutput, error) {
	_, err := us.FindByEmail(userDto.Email)
	if err == nil {
		return dto.CreateUserOutput{}, errors.New("user already exists")
	}

	user, err := us.userRepository.Create(User{
		Username: userDto.Username,
		Email:    userDto.Email,
		Password: userDto.Password,
	})

	if err != nil {
		return dto.CreateUserOutput{}, errors.New("failed to create user")
	}

	return dto.CreateUserOutput{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (us *UserService) Delete(email string) error {
	_, err := us.FindByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	err = us.userRepository.Delete(email)
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}

func (us *UserService) Login(userDto dto.LoginUserInput) (dto.LoginUserOutput, error) {
	user, err := us.userRepository.FindByEmail(userDto.Email)
	if err != nil {
		return dto.LoginUserOutput{}, errors.New("failed to find user")
	}

	if !user.ValidatePassword(userDto.Password) {
		return dto.LoginUserOutput{}, errors.New("invalid password")
	}
	token, err := us.AuthService.GenerateToken(user.Email)
	if err != nil {
		return dto.LoginUserOutput{}, errors.New("failed to generate token")
	}

	return dto.LoginUserOutput{
		Token: token,
	}, nil
}
