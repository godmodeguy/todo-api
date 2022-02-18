package service

import (
	"errors"
	"learn/todoapi/pkg/models"
	"learn/todoapi/pkg/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (as *AuthService) CreateUser(user models.User) (int, error) {
	return as.repo.CreateUser(user)
}

func (as *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := as.repo.GetUser(username, password)
	if err != nil || user.Username == "" {
		return "", errors.New("incorrect login or password")
	}
	logrus.Debugf("auth: %v:%v -> %+v", username, password, user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
		"username": user.Username,
		"issuedAt": time.Now().Unix(),
	})

	logrus.Debugf("jwt token: %v", token)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (as *AuthService) ParseToken(token string) (int, error) {
	// jwt.ParseWithClaims(token, &jwt.MapClaims{})
	return 0, nil
}