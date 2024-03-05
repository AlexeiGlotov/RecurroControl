package service

import (
	todo "RecurroControl"
	"RecurroControl/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Admission interface {
	CreateKey(userID int) (string, error)
	GetKey() ([]todo.RegAdmission, error)
}

type Service struct {
	Authorization
	Admission
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Admission:     NewAdmissionService(repos.Admission),
	}
}
