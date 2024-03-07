package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	todo "RecurroControl"
	"RecurroControl/internal/repository"
)

const (
	salt      = "dqwdqwdqwf12f432"
	signinKey = "dqwfqwf213122e1d121"
	tokenTTL  = 12 * time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	todo.User
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) ParseToken(accessToken string) (*todo.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signinKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of type *tokenClaims")
	}

	return &claims.User, nil
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, genereatePasswordHash(password))
	if err != nil {
		return "", fmt.Errorf("not user or bad password %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user,
	})
	return token.SignedString([]byte(signinKey))
}

func (s *AuthService) CreateUser(user todo.SignUpInput) (int, error) {
	user.Password = genereatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func genereatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CheckKeyAdmission(key string) (string, error) {
	return s.repo.CheckKeyAdmission(key)
}

func (s *AuthService) SetLoginAdmission(login, key string) error {
	return s.repo.SetLoginAdmission(login, key)
}
