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

func (l *LicenseKeysService) GetLicenseKeys(userID, limit, offset int, filter string) ([]models.LicenseKeys, error) {
	return l.repo.GetLicenseKeys(userID, limit, offset, filter)
}

func (l *LicenseKeysService) ResetHWID(id int) error {
	return l.repo.ResetHWID(id)
}

func (l *LicenseKeysService) Ban(id int) error {
	return l.repo.Ban(id)
}

func (l *LicenseKeysService) Unban(id int) error {
	return l.repo.Unban(id)
}

func (l *LicenseKeysService) Delete(id int) error {
	return l.repo.Delete(id)
}
