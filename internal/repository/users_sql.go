package repository

import (
	"database/sql"
	"fmt"

	todo "RecurroControl"
)

type UsersSql struct {
	db *sql.DB
}

func NewUsersSql(db *sql.DB) *UsersSql {
	return &UsersSql{db: db}
}

func (u *UsersSql) GetUsers() ([]todo.User, error) {
	users := []todo.User{}

	query := fmt.Sprintf("SELECT id,login,role,owner FROM %s", usersTable)
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
