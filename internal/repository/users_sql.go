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

func (u *UsersSql) Ban(userID int) error {
	query := fmt.Sprintf("UPDATE %s SET `banned` = 1 WHERE id = ?", usersTable)

	result, err := u.db.Exec(query, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}

	return nil
}

func (u *UsersSql) Unban(userID int) error {
	query := fmt.Sprintf("UPDATE %s SET `banned` = 0 WHERE id = ?", usersTable)

	result, err := u.db.Exec(query, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}

	return nil
}

func (u *UsersSql) Delete(userID int) error {
	query := fmt.Sprintf("UPDATE %s SET `is_deleted` = 1 WHERE id = ?", usersTable)

	result, err := u.db.Exec(query, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows updated")
	}

	return nil
}

func (u *UsersSql) GetUser(userID int) (*models.User, error) {
	user := models.User{}
	query := fmt.Sprintf("SELECT id,login,role,key_generated,key_activated,banned,owner,is_deleted FROM %s WHERE id = ?",
		usersTable)
	err := u.db.QueryRow(query, userID).Scan(&user.Id, &user.Login, &user.Role, &user.KeysGenerated,
		&user.KeysActivated, &user.Banned, &user.Owner, &user.IsDeleted)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UsersSql) GetUsers(userID int) ([]models.User, error) {
	users := []models.User{}
	var query string

	user, err := u.GetUser(userID)
	if err != nil {
		return nil, err
	}

	switch user.Role {
	case models.Admin:
		query = fmt.Sprintf("SELECT id,login,role,banned,owner,key_generated,key_activated FROM %s WHERE is_deleted = 0",
			usersTable)
	case models.Distributors, models.Reseller:
		query = fmt.Sprintf("SELECT id,login,role,banned,owner,key_generated,key_activated  FROM %s WHERE owner = '%s' and is_deleted = 0",
			usersTable,
			user.Login)
	default:
		return nil, errors.New("bad role")
	}

	row, err := u.db.Query(query)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := models.User{}
		err := row.Scan(&temp.Id,
			&temp.Login,
			&temp.Role,
			&temp.Banned,
			&temp.Owner,
			&temp.KeysGenerated,
			&temp.KeysActivated)
		if err != nil {
			return nil, err
		}
		users = append(users, temp)
	}

	if user.Role == models.Distributors {
		for _, x := range users {

			query = fmt.Sprintf("SELECT id,login,role,banned,owner,key_generated,key_activated  FROM %s WHERE owner = '%s' and is_deleted = 0",
				usersTable,
				x.Login)
			row, err := u.db.Query(query)

			if err != nil {
				return nil, err
			}

			for row.Next() {
				temp := models.User{}
				err := row.Scan(&temp.Id,
					&temp.Login,
					&temp.Role,
					&temp.Banned,
					&temp.Owner,
					&temp.KeysGenerated,
					&temp.KeysActivated)
				if err != nil {
					return nil, err
				}
				users = append(users, temp)
			}
		}
	}
	return users, nil
}
