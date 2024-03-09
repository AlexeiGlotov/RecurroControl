package repository

import "database/sql"

type LicenseKeysSql struct {
	db *sql.DB
}

func NewLicenseKeysSql(db *sql.DB) *LicenseKeysSql {
	return &LicenseKeysSql{db: db}
}

func (l *LicenseKeysSql) CreateLicenseKeys() {

}
