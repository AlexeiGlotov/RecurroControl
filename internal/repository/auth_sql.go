package repository

import (
	"database/sql"
	"fmt"

	"RecurroControl/models"
)

type AuthSql struct {
	db *sql.DB
}

func NewAuthSql(db *sql.DB) *AuthSql {
	return &AuthSql{db: db}
}

func (a *AuthSql) CreateUser(user models.SignUpInput) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (login ,password_hash,owner,role) values (?,?,?,?)", usersTable)
	row := a.db.QueryRow(query, user.Login, user.Password, user.Owner, user.Role)
	if row.Err() != nil {
		return 0, row.Err()
	}

	row = a.db.QueryRow(fmt.Sprintf("SELECT MAX(id) FROM %s;", usersTable))

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthSql) GetUser(username, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id,login,role,banned,owner FROM %s WHERE login=? AND password_hash=?", usersTable)

	row := a.db.QueryRow(query, username, password).Scan(&user.Id, &user.Login, &user.Role, &user.Banned, &user.Owner)

	return user, row
}

func (a *AuthSql) CheckAccessKey(key string) (*models.AccessKey, error) {

	var access_key models.AccessKey
	query := fmt.Sprintf("SELECT owner,role FROM %s WHERE access_key = ? and is_login is NULL", admissionTable)

	row := a.db.QueryRow(query, key)
	if err := row.Scan(&access_key.Owner, &access_key.Role); err != nil {
		return nil, err
	}

	return &access_key, nil
}

// Устанавливает логин , который использовал ключ
func (a *AuthSql) SetLoginAccessKey(login, key string) error {

	query := fmt.Sprintf("UPDATE %s SET `is_login`= ? WHERE access_key = ?", admissionTable)

	row := a.db.QueryRow(query, login, key)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
