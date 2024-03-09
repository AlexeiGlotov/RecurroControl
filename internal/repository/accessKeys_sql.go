package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"RecurroControl/models"
)

type AdmissionSql struct {
	db *sql.DB
}

func NewAdmissionSql(db *sql.DB) *AdmissionSql {
	return &AdmissionSql{db: db}
}

func generateUniqueKey() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

	seed := int64(time.Now().UnixNano())
	src := rand.NewSource(seed)
	r := rand.New(src)

	key := make([]byte, 16)
	for i := 0; i < 16; i++ {
		key[i] = charset[r.Intn(len(charset))]
	}

	return string(key)
}

func (a *AdmissionSql) CreateAccessKey(userID int, role string) (string, error) {
	var login string

	query := fmt.Sprintf("SELECT login FROM %s WHERE id=?", usersTable)
	row := a.db.QueryRow(query, userID)

	if row.Err() != nil {
		return "", row.Err()
	}

	err := row.Scan(&login)
	if err != nil {
		return "", err
	}

	key_gen := generateUniqueKey()
	query = fmt.Sprintf("INSERT INTO %s (access_key ,owner,role) values (?,?,?)", admissionTable)
	row = a.db.QueryRow(query, key_gen, login, role)

	if row.Err() != nil {
		return "", row.Err()
	}

	return key_gen, nil
}

func (a *AdmissionSql) GetAccessKey(login, role string) ([]models.AccessKey, error) {

	ra := []models.AccessKey{}
	var query string

	switch role {
	case models.Admin:
		query = fmt.Sprintf("SELECT * FROM %s", admissionTable)
	case models.Distributors:
		query = fmt.Sprintf("SELECT * FROM %s WHERE owner = '%s'", admissionTable, login)
	default:
		return nil, errors.New("bad role")
	}

	row, err := a.db.Query(query)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := models.AccessKey{}
		err := row.Scan(&temp.Id, &temp.AccessKey, &temp.Owner, &temp.Role, &temp.IsLogin)
		if err != nil {
			return nil, err
		}
		ra = append(ra, temp)
	}

	return ra, nil
}
