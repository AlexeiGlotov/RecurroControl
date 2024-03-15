package repository

import (
	"database/sql"
	"errors"
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

func buildQuery(filters []string, value string) string {

	var conditions []string
	for _, iter := range filters {
		conditions = append(conditions, fmt.Sprintf("%s LIKE '%%%s%%'", iter, value))
	}

	whereClause := strings.Join(conditions, " OR ")
	return fmt.Sprintf("%s", whereClause)
}

func (l *LicenseKeysSql) GetLicenseKeys(
	login,
	role string,
	limit, offset int,
	filter string,
) ([]models.LicenseKeys, error) {
	var cheats []models.LicenseKeys

	or := buildQuery([]string{
		"license_key",
		"cheat",
		"ttl_cheat",
		"holder",
		"creator",
		"date_creation",
		"date_activation",
		"hwid",
		"hwidk",
	}, filter)

	var query string
	switch role {
	case models.Admin:
		query = fmt.Sprintf("SELECT id, license_key, cheat, ttl_cheat, holder, creator, date_creation, date_activation, hwid, hwidk, banned, is_deleted FROM %s WHERE %s LIMIT ? OFFSET ?",
			licenseKeysTable, or)
	case models.Distributors, models.Reseller:
		query = fmt.Sprintf("SELECT id, license_key, cheat, ttl_cheat, holder, creator, date_creation, date_activation, hwid, hwidk, banned, is_deleted FROM %s WHERE (holder = '%s' or creator = '%s') AND (%s)  LIMIT ? OFFSET ?",
			licenseKeysTable, login, login, or)
	case models.Salesman:
		query = fmt.Sprintf("SELECT id, license_key, cheat, ttl_cheat, holder, creator, date_creation, date_activation, hwid, hwidk, banned, is_deleted FROM %s WHERE (holder = '%s') AND (%s)  LIMIT ? OFFSET ?",
			licenseKeysTable, login, or)
	}

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
			&temp.DateCreation,
			&temp.DateActivation,
			&temp.HWID,
			&temp.HWIDK,
			&temp.Banned,
			&temp.IsDeleted); err != nil {
			return nil, err
		}
		cheats = append(cheats, temp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cheats, nil
}

func (l *LicenseKeysSql) Ban(id int) error {
	query := fmt.Sprintf("UPDATE %s SET `banned` = 1 WHERE id = ?", licenseKeysTable)

	result, err := l.db.Exec(query, id)
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

func (l *LicenseKeysSql) Unban(id int) error {
	query := fmt.Sprintf("UPDATE %s SET `banned` = 0 WHERE id = ?", licenseKeysTable)

	result, err := l.db.Exec(query, id)
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

func (l *LicenseKeysSql) Delete(id int) error {
	query := fmt.Sprintf("UPDATE %s SET `is_deleted` = 1 WHERE id = ?", licenseKeysTable)

	result, err := l.db.Exec(query, id)
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

func (l *LicenseKeysSql) ResetHWID(id int) error {
	query := fmt.Sprintf("UPDATE %s SET `hwid` = null , hwidk = null WHERE id = ?", licenseKeysTable)

	result, err := l.db.Exec(query, id)
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
