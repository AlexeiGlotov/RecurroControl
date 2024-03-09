package repository

import (
	"database/sql"
	"errors"
	"fmt"

	todo "RecurroControl"
)

type UsersSql struct {
	db *sql.DB
}

func NewUsersSql(db *sql.DB) *UsersSql {
	return &UsersSql{db: db}
}

func (u *UsersSql) GetUserStruct(userID int) (*todo.User, error) {
	user := todo.User{}
	query := fmt.Sprintf("SELECT id,login,role,owner FROM %s WHERE id = ?", usersTable)
	err := u.db.QueryRow(query, userID).Scan(&user.Id, &user.Login, &user.Role, &user.Owner)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersSql) GetUserLoginsAndRole(userID int) ([]todo.User, error) {
	users := []todo.User{}
	var query string

	user, err := u.GetUserStruct(userID)
	if err != nil {
		return nil, err
	}

	switch user.Role {
	case todo.Admin:
		query = fmt.Sprintf("SELECT id,login,role,owner FROM %s", usersTable)
	case todo.Moder:
		query = fmt.Sprintf("SELECT id,login,role,owner FROM %s WHERE owner = '%s' or login ='%s'",
			usersTable,
			user.Login,
			user.Login)
	case todo.Seller:
		query = fmt.Sprintf("SELECT id,login,role,owner FROM %s WHERE login = '%s'", usersTable, user.Login)
	case todo.Reseller:
		query = fmt.Sprintf("SELECT id,login,role,owner FROM %s WHERE login = '%s'", usersTable, user.Login)
	default:
		return nil, errors.New("bad role")
	}

	row, err := u.db.Query(query)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := todo.User{}
		err := row.Scan(&temp.Id, &temp.Login, &temp.Role, &temp.Owner)
		if err != nil {
			return nil, err
		}
		users = append(users, temp)
	}

	return users, nil
}
