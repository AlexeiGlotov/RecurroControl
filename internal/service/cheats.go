package service

import (
	"RecurroControl/internal/repository"
	"RecurroControl/models"
)

type CheatService struct {
	repo repository.Cheats
}

func NewCheatService(repo repository.Cheats) *CheatService {
	return &CheatService{repo: repo}
}

func (c *CheatService) GetCheats(role string) ([]models.Cheats, error) {
	return c.repo.GetCheats(role)
}

func (c *CheatService) CreateCheats(cheat *models.Cheats) (int, error) {
	return c.repo.CreateCheats(cheat)
}
