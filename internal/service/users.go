package service

import (
	todo "RecurroControl"
	"RecurroControl/internal/repository"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (u *UsersService) GetUsers() ([]todo.User, error) {
	return u.repo.GetUsers()
}
