package service

import (
	"RecurroControl/internal/repository"
	"RecurroControl/models"
)

type UsersService struct {
	repo repository.Users
}

func NewUsersService(repo repository.Users) *UsersService {
	return &UsersService{repo: repo}
}

func (u *UsersService) GetUserLoginsAndRole(userID int) ([]models.User, error) {
	return u.repo.GetUserLoginsAndRole(userID)
}

func (u *UsersService) GetUserStruct(userID int) (*models.User, error) {
	return u.repo.GetUserStruct(userID)
}
