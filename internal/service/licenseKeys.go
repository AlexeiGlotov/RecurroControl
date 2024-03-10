package service

import (
	"RecurroControl/internal/repository"
	"RecurroControl/models"
)

type LicenseKeysService struct {
	repo repository.LicenseKeys
}

func NewLicenseKeysService(repo repository.LicenseKeys) *LicenseKeysService {
	return &LicenseKeysService{repo: repo}
}

func (l *LicenseKeysService) CreateLicenseKeys(keys []models.LicenseKeys) error {
	return l.repo.CreateLicenseKeys(keys)
}
