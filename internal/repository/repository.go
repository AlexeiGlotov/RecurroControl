package repository

import (
	"database/sql"

	"RecurroControl/models"
)

const (
	usersTable     = "users"
	admissionTable = "access_key"
	cheatTable     = "cheats"
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
	GetUserLoginsAndRole(userID int) ([]models.User, error)
	GetUserStruct(userID int) (*models.User, error)
}

type Cheats interface {
	GetCheats() ([]models.Cheats, error)
}

type Repository struct {
	Authorization
	AccessKeys
	Cheats
	Users
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		AccessKeys:    NewAdmissionSql(db),
		Cheats:        NewCheatSql(db),
		Users:         NewUsersSql(db),
	}
}
