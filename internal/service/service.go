package service

import (
	todo "RecurroControl"
	"RecurroControl/internal/repository"
)

type Authorization interface {
	CreateUser(user todo.SignUpInput) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (*todo.User, error)
	CheckKeyAdmission(key string) (string, error)
	SetLoginAdmission(login, key string) error
}

type Admission interface {
	CreateKey(userID int) (string, error)
	GetKey() ([]todo.RegAdmission, error)
}

type Cheats interface {
	GetCheats() ([]todo.StCheats, error)
}

type Users interface {
	GetUsers() ([]todo.User, error)
}
type Service struct {
	Authorization
	Admission
	Cheats
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Admission:     NewAdmissionService(repos.Admission),
		Cheats:        NewCheatService(repos.Cheats),
		Users:         NewUsersService(repos.Users),
	}
}
