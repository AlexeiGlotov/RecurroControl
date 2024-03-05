package service

import (
	todo "RecurroControl"
	"RecurroControl/pkg/repository"
)

type AdmissionService struct {
	repo repository.Admission
}

func NewAdmissionService(repo repository.Admission) *AdmissionService {
	return &AdmissionService{repo: repo}
}

func (s *AdmissionService) CreateKey(userID int) (string, error) {
	return s.repo.CreateKey(userID)
}

func (s *AdmissionService) GetKey() ([]todo.RegAdmission, error) {
	return s.repo.GetKey()
}
