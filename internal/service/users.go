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

func (u *UsersService) GetUsers(userID int) ([]models.User, error) {
	return u.repo.GetUsers(userID)
}

func (u *UsersService) GetUser(userID int) (*models.User, error) {
	return u.repo.GetUser(userID)
}

func (u *UsersService) Ban(userID int) error {
	return u.repo.Ban(userID)
}

func (u *UsersService) Unban(userID int) error {
	return u.repo.Unban(userID)
}

func (u *UsersService) Delete(userID int) error {
	return u.repo.Delete(userID)
}
