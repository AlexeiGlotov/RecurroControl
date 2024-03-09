package service

import (
	"RecurroControl/internal/repository"
	"RecurroControl/models"
)

type Authorization interface {
	CreateUser(user models.SignUpInput) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	CheckAccessKey(key string) (*models.AccessKey, error)
	SetLoginAccessKey(login, key string) error
}

type AccessKeys interface {
	CreateAccessKey(userID int, role string) (string, error)
	GetAccessKey(login, role string) ([]models.AccessKey, error)
}

type Cheats interface {
	GetCheats(role string) ([]models.Cheats, error)
	CreateCheats(cheat *models.Cheats) (int, error)
}

type Users interface {
	GetUserLoginsAndRole(userID int) ([]models.User, error)
	GetUserStruct(userID int) (*models.User, error)
}
type Service struct {
	Authorization
	AccessKeys
	Cheats
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		AccessKeys:    NewAdmissionService(repos.AccessKeys),
		Cheats:        NewCheatService(repos.Cheats),
		Users:         NewUsersService(repos.Users),
	}
}
