package repository

import (
	"database/sql"

	todo "RecurroControl"
)

const (
	usersTable     = "users"
	admissionTable = "reg_admission"
)

type Authorization interface {
	CreateUser(user todo.SignUpInput) (int, error)
	GetUser(username, password string) (todo.User, error)
	CheckKeyAdmission(key string) (string, error)
}

type Admission interface {
	CreateKey(userID int) (string, error)
	GetKey() ([]todo.RegAdmission, error)
}

type Repository struct {
	Authorization
	Admission
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSql(db),
		Admission:     NewAdmissionSql(db),
	}
}
