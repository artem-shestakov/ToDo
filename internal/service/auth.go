package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/artem-shestakov/to-do/internal/models"
	"github.com/artem-shestakov/to-do/internal/repository"
	"github.com/golang-jwt/jwt"
)

const (
	salt       = "qwerty"
	signingKey = "ytrewq"
	jwtTTL     = 12 * time.Hour
)

type AuthService struct {
	repo *repository.Repository
}

type customClaims struct {
	jwt.StandardClaims
	UserID    int
	UserEmail string
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

func (s *AuthService) GetUser(email, password string) (models.User, error) {
	return s.repo.Auth.GetUser(email, password)
}

func (s *AuthService) passwordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.GetUser(email, s.passwordHash(password))
	if err != nil {
		return "", err
	}
	claims := customClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
		user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}
