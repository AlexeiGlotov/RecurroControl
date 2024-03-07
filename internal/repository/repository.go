package repository

import (
	"database/sql"

	todo "RecurroControl"
)

const (
	usersTable     = "users"
	admissionTable = "reg_admission"
	cheatTable     = "cheats"
)

type Authorization interface {
	CreateUser(user todo.SignUpInput) (int, error)
	GetUser(username, password string) (todo.User, error)
	CheckKeyAdmission(key string) (string, error)
	SetLoginAdmission(login, key string) error
}

type Admission interface {
	CreateKey(userID int) (string, error)
	GetKey() ([]todo.RegAdmission, error)
}

type Users interface {
	GetUsers() ([]todo.User, error)
}

type Cheats interface {
	GetCheats() ([]todo.StCheats, error)
}

type Repository struct {
	Authorization
	Admission
	Cheats
	Users
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		Admission:     NewAdmissionSql(db),
		Cheats:        NewCheatSql(db),
		Users:         NewUsersSql(db),
	}
}
