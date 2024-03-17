package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"RecurroControl/internal/repository"
	"RecurroControl/models"
)

const (
	tokenTTL = time.Hour * 6
)

type AuthService struct {
	repo         repository.Authorization
	SaltPassword string
	SaltJWT      string
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int
	Role   string
}

func NewAuthService(repo repository.Authorization, saltPassword, saltJWT string) *AuthService {
	return &AuthService{repo: repo, SaltPassword: saltPassword, SaltJWT: saltJWT}
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.SaltJWT), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, genereatePasswordHash(password, s.SaltPassword))
	if err != nil {
		return "", fmt.Errorf("not user or bad password %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Role,
	})
	return token.SignedString([]byte(s.SaltJWT))
}

func (s *AuthService) CreateUser(user models.SignUpInput) (int, error) {
	user.Password = genereatePasswordHash(user.Password, s.SaltPassword)
	return s.repo.CreateUser(user)
}
func genereatePasswordHash(password string, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CheckAccessKey(key string) (*models.AccessKey, error) {
	return s.repo.CheckAccessKey(key)
}

func (s *AuthService) SetLoginAccessKey(login, key string) error {
	return s.repo.SetLoginAccessKey(login, key)
}
