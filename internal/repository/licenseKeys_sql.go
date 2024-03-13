package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"RecurroControl/models"
)

type LicenseKeysSql struct {
	db *sql.DB
}

func NewLicenseKeysSql(db *sql.DB) *LicenseKeysSql {
	return &LicenseKeysSql{db: db}
}

func (l *LicenseKeysSql) CreateLicenseKeys(keys []models.LicenseKeys) error {

	var valueStrings []string
	var valueArgs []interface{}

	for _, key := range keys {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, key.LicenseKeys, key.Cheat, key.TTLCheat, key.Holder, key.Creator)
	}

	stmt := fmt.Sprintf("INSERT INTO %s (license_key, cheat, ttl_cheat, holder, creator) VALUES %s",
		licenseKeysTable, strings.Join(valueStrings, ","))

	_, err := l.db.Exec(stmt, valueArgs...)
	if err != nil {
		return err
	}

	return nil
}

func (l *LicenseKeysSql) GetLicenseKeys(userID, limit, offset int) ([]models.LicenseKeys, error) {
	var cheats []models.LicenseKeys

	query := fmt.Sprintf("SELECT id, license_key, cheat, ttl_cheat, holder, creator, date_creation FROM %s LIMIT ? OFFSET ?",
		licenseKeysTable)
	rows, err := l.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var temp models.LicenseKeys
		if err := rows.Scan(&temp.Id,
			&temp.LicenseKeys,
			&temp.Cheat,
			&temp.TTLCheat,
			&temp.Holder,
			&temp.Creator,
			&temp.DateCreation); err != nil {
			return nil, err
		}
		cheats = append(cheats, temp)
	}

	// Проверяем на наличие ошибок при итерации по результатам
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cheats, nil
}
