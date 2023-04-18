package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/artem-shestakov/to-do/internal/repository"
)

const salt = "qwerty"

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = s.passwordHash(user.Password)
	return s.repo.Auth.CreateUser(user)
}

func (s *AuthService) passwordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
