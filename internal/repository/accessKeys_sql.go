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
	const charset = "ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnopqrstuvwxyz123456789"

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

	var query string

	switch role {
	case models.Admin:
		query = "SELECT * FROM " + admissionTable
		return a.executeAccessKeyQuery(query, nil)
	case models.Distributors:
		query = "SELECT * FROM " + admissionTable + " WHERE owner = ?"
		return a.executeAccessKeyQueryForDistributor(query, login)
	case models.Reseller:
		query = "SELECT * FROM " + admissionTable + " WHERE owner = ?"
		return a.executeAccessKeyQuery(query, []interface{}{login})
	default:
		return nil, errors.New("bad role")
	}
}

func (a *AdmissionSql) executeAccessKeyQuery(query string, args []interface{}) ([]models.AccessKey, error) {
	var ra []models.AccessKey
	rows, err := a.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := models.AccessKey{}
		if err := rows.Scan(&temp.Id, &temp.AccessKey, &temp.Owner, &temp.Role, &temp.IsLogin); err != nil {
			return nil, err
		}
		ra = append(ra, temp)
	}

	return ra, rows.Err()
}

func (a *AdmissionSql) executeAccessKeyQueryForDistributor(query string, login string) ([]models.AccessKey, error) {
	keys, err := a.executeAccessKeyQuery(query, []interface{}{login})
	if err != nil {
		return nil, err
	}

	var allKeys []models.AccessKey
	allKeys = append(allKeys, keys...)

	for _, key := range keys {
		if key.IsLogin != nil && *key.IsLogin != login {
			moreKeys, err := a.executeAccessKeyQuery(query, []interface{}{*key.IsLogin})
			if err != nil {
				return nil, err
			}
			allKeys = append(allKeys, moreKeys...)
			login = *key.IsLogin
		}
	}

	return allKeys, nil
}
