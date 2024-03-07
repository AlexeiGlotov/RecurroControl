package service

import (
	todo "RecurroControl"
	"RecurroControl/internal/repository"
)

type CheatService struct {
	repo repository.Cheats
}

func NewCheatService(repo repository.Cheats) *CheatService {
	return &CheatService{repo: repo}
}

func (c *CheatService) GetCheats() ([]todo.StCheats, error) {
	return c.repo.GetCheats()
}
