package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/dto"
	"github.com/ivanovamir/Clean-architectur-rest-API/internal/repository"
	"os"
	"time"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repo: repo,
	}
}

type CustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int) string {
	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(4320 * time.Hour)), //TODO mb change expires time
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNED_KEY")))
	return tokenString
}

func (s *UserService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SIGNED_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *CustomClaims")
	}

	return claims.UserId, nil
}

func (s *UserService) RegisterUser(userDTO *dto.RegisterUser) (string, error) {
	newUserId, err := s.repo.RegisterUser(userDTO)
	if err != nil {
		return "", err
	} else {
		token := GenerateToken(newUserId)
		return token, nil
	}
}
