package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/matheusvidal21/smart-news-fetcher/internal/dto"
	"github.com/matheusvidal21/smart-news-fetcher/internal/models"
)

type UserRepositoryInterface interface {
	FindByEmail(email string) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Delete(email string) error
	Update(user models.User) (*models.User, error)
	FindById(id int) (*models.User, error)
}

type UserServiceInterface interface {
	FindByEmail(email string) (dto.FindUserByEmailOutput, error)
	Create(userDto dto.CreateUserInput) (dto.CreateUserOutput, error)
	Delete(email string) error
	Login(userDto dto.LoginUserInput) (dto.LoginUserOutput, error)
	UpdatePassword(userDto dto.UpdateUserPasswordInput) error
	FindById(id int) (dto.FindUserByIdOutput, error)
}

type UserHandlerInterface interface {
	CreateUser(c *gin.Context)
	FindByEmail(c *gin.Context)
	DeleteUser(c *gin.Context)
	Login(c *gin.Context)
	UpdatePassword(c *gin.Context)
	FindById(c *gin.Context)
}
