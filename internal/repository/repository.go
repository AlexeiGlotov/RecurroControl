package repository

import (
	"database/sql"
	"time"

	"RecurroControl/models"
)

const (
	usersTable       = "users"
	admissionTable   = "access_key"
	cheatTable       = "cheats"
	licenseKeysTable = "license_keys"
)

type Authorization interface {
	CreateUser(user models.SignUpInput) (int, error)
	GetUser(username, password string) (models.User, error)
	CheckAccessKey(key string) (*models.AccessKey, error)
	SetLoginAccessKey(login, key string) error
}

type AccessKeys interface {
	CreateAccessKey(userID int, role string) (string, error)
	GetAccessKey(login, role string) ([]models.AccessKey, error)
}

type Users interface {
	GetUsers(userID int) ([]models.User, error)
	GetUser(userID int) (*models.User, error)
	Ban(userID int) error
	Unban(userID int) error
	Delete(userID int) error
	GetUsersAll() ([]models.User, error)
}

type Cheats interface {
	GetCheats(role string) ([]models.Cheats, error)
	CreateCheat(cheat *models.Cheats) (int, error)
	UpdateCheat(cheat *models.Cheats) error
}

type LicenseKeys interface {
	CreateLicenseKeys(keys []models.LicenseKeys) error
	GetLicenseKeys(limit, offset int, filter string) ([]models.LicenseKeys, error)
	GetCustomLicenseKeys(login string, date time.Time) (*models.InfoKeyDashboard, error)
	Delete(id int) error
	Ban(id int) error
	BanOfDate(login string, date time.Time) error
	Unban(id int) error
	ResetHWID(id int) error
}

type Repository struct {
	Authorization
	AccessKeys
	Cheats
	Users
	LicenseKeys
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		AccessKeys:    NewAdmissionSql(db),
		Cheats:        NewCheatSql(db),
		Users:         NewUsersSql(db),
		LicenseKeys:   NewLicenseKeysSql(db),
	}
}
