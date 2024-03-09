package service

import (
	"RecurroControl/internal/repository"
	"RecurroControl/models"
)

type AdmissionService struct {
	repo repository.AccessKeys
}

func NewAdmissionService(repo repository.AccessKeys) *AdmissionService {
	return &AdmissionService{repo: repo}
}

func (s *AdmissionService) CreateAccessKey(userID int, role string) (string, error) {
	return s.repo.CreateAccessKey(userID, role)
}

func (s *AdmissionService) GetAccessKey(login, role string) ([]models.AccessKey, error) {
	return s.repo.GetAccessKey(login, role)
}
