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

func (u *UsersService) GetUserLoginsAndRole(userID int) ([]todo.User, error) {
	return u.repo.GetUserLoginsAndRole(userID)
}

func (u *UsersService) GetUserStruct(userID int) (*todo.User, error) {
	return u.repo.GetUserStruct(userID)
}
