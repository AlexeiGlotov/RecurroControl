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

func (l *LicenseKeysService) GetLicenseKeys(userID, limit, offset int) ([]models.LicenseKeys, error) {
	return l.repo.GetLicenseKeys(userID, limit, offset)
}
