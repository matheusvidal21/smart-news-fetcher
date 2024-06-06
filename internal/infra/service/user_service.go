package service

import (
	"errors"
	"github.com/matheusvidal21/smart-news-fetcher/internal/auth"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/email"
	"github.com/matheusvidal21/smart-news-fetcher/internal/interfaces"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
)

type UserService struct {
	userRepository interfaces.UserRepositoryInterface
	EmailService   interfaces.EmailService
	AuthService    auth.JWTServiceInterface
}

func NewUserService(userRepository interfaces.UserRepositoryInterface, emailService interfaces.EmailService, authService auth.JWTServiceInterface) *UserService {
	return &UserService{
		userRepository: userRepository,
		EmailService:   emailService,
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

	user, err := us.userRepository.Create(models.User{
		Username: userDto.Username,
		Email:    userDto.Email,
		Password: userDto.Password,
	})

	if err != nil {
		return dto.CreateUserOutput{}, errors.New("failed to create user: " + err.Error())
	}

	message := email.Message{
		ToEmail:          userDto.Email,
		Subject:          "Welcome to Smart News Fetcher",
		PlainTextContent: "Welcome to Smart News Fetcher, " + userDto.Username + "!",
		HtmlContent:      "<p> Welcome to Smart News Fetcher, <b>" + userDto.Username + "! </b> </p>",
	}

	err = us.EmailService.Send(message)

	if err != nil {
		return dto.CreateUserOutput{}, errors.New("User created but failed to send email: " + err.Error())
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
		return errors.New("failed to delete user: " + err.Error())
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
		return dto.LoginUserOutput{}, errors.New("failed to generate token: " + err.Error())
	}

	return dto.LoginUserOutput{
		Token: token,
	}, nil
}

func (us *UserService) UpdatePassword(userDto dto.UpdateUserPasswordInput) error {
	user, err := us.userRepository.FindByEmail(userDto.Email)
	if err != nil {
		return errors.New("failed to find user")
	}

	ok := user.ValidatePassword(userDto.OldPassword)
	if !ok {
		return errors.New("invalid password")
	}

	user.Password = userDto.NewPassword
	_, err = us.userRepository.Update(*user)

	if err != nil {
		return errors.New("failed to update password: " + err.Error())
	}
	return nil
}

func (us *UserService) FindById(id int) (dto.FindUserByIdOutput, error) {
	user, err := us.userRepository.FindById(id)
	if err != nil {
		return dto.FindUserByIdOutput{}, errors.New("failed to find user")
	}
	return dto.FindUserByIdOutput{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
