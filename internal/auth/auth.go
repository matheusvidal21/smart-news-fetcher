package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTServiceInterface interface {
	GenerateToken(email string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JWTService struct {
	secretKey         string
	expirationMinutes int
}

func NewJWTService(secretKey string, expirationMinutes int) *JWTService {
	return &JWTService{
		secretKey:         secretKey,
		expirationMinutes: expirationMinutes,
	}
}

func (s *JWTService) GenerateToken(email string) (string, error) {
	claims := &authCustomClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.expirationMinutes) * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *JWTService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(encodedToken, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.secretKey), nil
	})

	if _, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
		return token, nil
	}
	return nil, err
}
