package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"RecurroControl/models"
)

type UsersSql struct {
	db *sql.DB
}

func NewUsersSql(db *sql.DB) *UsersSql {
	return &UsersSql{db: db}
}

func (u *UsersSql) GetUserStruct(userID int) (*models.User, error) {
	user := models.User{}
	query := fmt.Sprintf("SELECT id,login,role,owner FROM %s WHERE id = ?", usersTable)
	err := u.db.QueryRow(query, userID).Scan(&user.Id, &user.Login, &user.Role, &user.Owner)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersSql) GetUserLoginsAndRole(userID int) ([]models.User, error) {
	users := []models.User{}
	var query string

	user, err := u.GetUserStruct(userID)
	if err != nil {
		return nil, err
	}

	switch user.Role {
	case models.Admin:
		query = fmt.Sprintf("SELECT id,login,role,owner FROM %s", usersTable)
	case models.Distributors:
		query = fmt.Sprintf("SELECT id,login,role,owner FROM %s WHERE owner = '%s' or login ='%s'",
			usersTable,
			user.Login,
			user.Login)
	case models.Reseller:
		query = fmt.Sprintf("SELECT id,login,role,owner FROM %s WHERE login = '%s'", usersTable, user.Login)
	case models.Salesman:
		query = fmt.Sprintf("SELECT id,login,role,owner FROM %s WHERE login = '%s'", usersTable, user.Login)
	default:
		return nil, errors.New("bad role")
	}

	row, err := u.db.Query(query)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := models.User{}
		err := row.Scan(&temp.Id, &temp.Login, &temp.Role, &temp.Owner)
		if err != nil {
			return nil, err
		}
		users = append(users, temp)
	}

	return users, nil
}
