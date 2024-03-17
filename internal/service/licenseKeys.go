package service

import (
	"time"

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

func (l *LicenseKeysService) GetLicenseKeys(login, role string, hier []string, limit, offset int, filter string) ([]models.LicenseKeys, error) {

	var returnData []models.LicenseKeys
	licenseKeys, err := l.repo.GetLicenseKeys(limit, offset, filter)
	if err != nil {
		return nil, err
	}

	if role == models.Admin {
		for _, item := range licenseKeys {
			returnData = append(returnData, item)
		}
		return returnData, nil
	}

	if role == models.Salesman {
		for _, item := range licenseKeys {
			if item.Holder == login {
				returnData = append(returnData, item)
			}
		}
		return returnData, nil
	}

	if role == models.Reseller {
		for _, item := range licenseKeys {
			if item.Holder == login || item.Creator == login {
				returnData = append(returnData, item)
			}
		}
		return returnData, nil
	}

	hier = append(hier, login)
	if role == models.Distributors {
		for _, item := range licenseKeys {
			for _, user := range hier {
				if item.Holder == user || item.Creator == user {
					returnData = append(returnData, item)
				}
			}
		}
		return returnData, nil
	}

	return nil, err
}

func (l *LicenseKeysService) ResetHWID(id int) error {
	return l.repo.ResetHWID(id)
}

func (l *LicenseKeysService) Ban(id int) error {
	return l.repo.Ban(id)
}

func (l *LicenseKeysService) BanOfDate(login string, date time.Time) error {
	return l.repo.BanOfDate(login, date)
}

func (l *LicenseKeysService) Unban(id int) error {
	return l.repo.Unban(id)
}

func (l *LicenseKeysService) Delete(id int) error {
	return l.repo.Delete(id)
}

func (l *LicenseKeysService) GetCustomLicenseKeys(login string, date time.Time) (*models.InfoKeyDashboard, error) {
	return l.repo.GetCustomLicenseKeys(login, date)
}
