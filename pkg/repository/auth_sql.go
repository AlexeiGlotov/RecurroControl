package repository

import (
	"database/sql"
	"fmt"

	todo "RecurroControl"
)

type AuthSql struct {
	db *sql.DB
}

func NewAuthSql(db *sql.DB) *AuthSql {
	return &AuthSql{db: db}
}

func (a *AuthSql) CreateUser(user todo.SignUpInput) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (login ,password,owner) values (?,?,?)", usersTable)
	row := a.db.QueryRow(query, user.Login, user.Password, user.Owner)
	if row.Err() != nil {
		return 0, row.Err()
	}

	row = a.db.QueryRow(fmt.Sprintf("SELECT MAX(id) FROM %s;", usersTable))

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthSql) GetUser(username, password string) (todo.User, error) {
	var user todo.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE login=? AND password=?", usersTable)

	row := a.db.QueryRow(query, username, password).Scan(&user.Id)

	return user, row
}

func (a *AuthSql) CheckKeyAdmission(key string) (string, error) {

	var owner string
	query := fmt.Sprintf("SELECT owner FROM %s WHERE access_key = ? and isLogin is NULL", admissionTable)

	row := a.db.QueryRow(query, key)
	if err := row.Scan(&owner); err != nil {
		return "", err
	}

	return owner, nil
}
