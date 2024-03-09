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

func (c *CheatService) GetCheats() ([]models.Cheats, error) {
	return c.repo.GetCheats()
}
