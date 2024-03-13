package repository

import (
	"database/sql"
	"fmt"

	"RecurroControl/models"
)

type LicenseKeysSql struct {
	db *sql.DB
}

func NewLicenseKeysSql(db *sql.DB) *LicenseKeysSql {
	return &LicenseKeysSql{db: db}
}

func (l *LicenseKeysSql) CreateLicenseKeys(keys []models.LicenseKeys) error {

	for _, key := range keys {
		query := fmt.Sprintf("INSERT INTO %s (license_key,cheat,ttl_cheat,holder,creator) values (?,?,?,?,?)",
			licenseKeysTable)
		row := l.db.QueryRow(query, key.LicenseKeys, key.Cheat, key.TTLCheat, key.Holder, key.Creator)
		if row.Err() != nil {
			return row.Err()
		}
	}
	return nil
}
