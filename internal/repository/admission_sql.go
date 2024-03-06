package repository

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	todo "RecurroControl"
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

func (a *AdmissionSql) CreateKey(userID int) (string, error) {
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
	query = fmt.Sprintf("INSERT INTO %s (access_key ,owner) values (?,?)", admissionTable)
	row = a.db.QueryRow(query, key_gen, login)

	if row.Err() != nil {
		return "", row.Err()
	}

	return key_gen, nil
}

func (a *AdmissionSql) GetKey() ([]todo.RegAdmission, error) {

	ra := []todo.RegAdmission{}

	query := fmt.Sprintf("SELECT * FROM %s", admissionTable)
	row, err := a.db.Query(query)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := todo.RegAdmission{}
		err := row.Scan(&temp.Id, &temp.Access_key, &temp.Owner, &temp.IsLogin)
		if err != nil {
			return nil, err
		}
		ra = append(ra, temp)
	}

	return ra, nil
}
