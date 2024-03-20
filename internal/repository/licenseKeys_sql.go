package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"RecurroControl/models"
)

type LicenseKeysSql struct {
	db *sql.DB
}

func NewLicenseKeysSql(db *sql.DB) *LicenseKeysSql {
	return &LicenseKeysSql{db: db}
}

func (l *LicenseKeysSql) CreateLicenseKeys(keys []models.LicenseKeys) error {
	// Начало транзакции
	tx, err := l.db.Begin()
	if err != nil {
		return err
	}

	var valueStrings []string
	var valueArgs []interface{}
	var login string

	for _, key := range keys {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, key.LicenseKeys, key.Cheat, key.TTLCheat, key.Holder, key.Creator, key.Notes)
		login = key.Holder
	}

	insert := fmt.Sprintf("INSERT INTO %s (license_key, cheat, ttl_cheat, holder, creator, notes) VALUES %s",
		licenseKeysTable, strings.Join(valueStrings, ","))

	update := fmt.Sprintf("UPDATE %s SET key_generated = key_generated + %d WHERE login = '%s'",
		usersTable, len(keys), login)

	// Выполнение обновления
	if _, err := tx.Exec(update); err != nil {
		tx.Rollback() // Откат в случае ошибки
		return err
	}

	// Выполнение вставки
	if _, err := tx.Exec(insert, valueArgs...); err != nil {
		tx.Rollback() // Откат в случае ошибки
		return err
	}

	// Подтверждение транзакции
	return tx.Commit()
}

func (l *LicenseKeysSql) GetLicenseKeys(limit, offset int, filter string) ([]models.LicenseKeys, error) {
	var cheats []models.LicenseKeys
	var query string
	var rows *sql.Rows
	var err error

	if filter == "" {
		query = fmt.Sprintf(`
			SELECT id, license_key, cheat, ttl_cheat, holder, creator, date_creation, date_activation,hwid, hwidk, 
			       banned, is_deleted,notes 
			FROM license_keys ORDER BY date_creation DESC LIMIT ? OFFSET ?`)
	} else {
		query = fmt.Sprintf(`
			SELECT id, license_key, cheat, ttl_cheat, holder, creator, date_creation, date_activation,hwid, hwidk, 
			       banned, is_deleted,notes 
			FROM license_keys 
			WHERE %s ORDER BY date_creation DESC LIMIT ? OFFSET ?`, filter)
	}

	rows, err = l.db.Query(query, limit, offset)

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
			&temp.IsDeleted,
			&temp.Notes); err != nil {
			return nil, err
		}
		cheats = append(cheats, temp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cheats, nil
}

func (l *LicenseKeysSql) GetCustomLicenseKeys(login string, date time.Time) (*models.InfoKeyDashboard, error) {

	var dashboard models.InfoKeyDashboard

	query := fmt.Sprintf(`
		SELECT COUNT(*) AS created, COUNT(date_activation) AS activated ,COALESCE(SUM(banned), 0) AS banned
		FROM 
			license_keys 
		WHERE 
			(holder = '%s' OR creator = '%s') 
			AND DATE(date_creation) = '%s';
	`, login, login, date.Format("2006-01-02"))

	row := l.db.QueryRow(query)
	err := row.Scan(&dashboard.CountAll, &dashboard.CountActive, &dashboard.CountBan)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &dashboard, nil
}

func (l *LicenseKeysSql) BanOfDate(login string, date time.Time) error {
	query := fmt.Sprintf(`
		UPDATE license_keys 
		SET banned = 1
		WHERE 
    	(holder = '%s' OR creator = '%s') 
    	AND DATE(date_creation) = '%s'`, login, login, date.Format("2006-01-02"))

	result, err := l.db.Exec(query)
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

func (l *LicenseKeysSql) Ban(id int) error {
	query := fmt.Sprintf("UPDATE %s "+
		"SET `banned` = 1 WHERE id = ?", licenseKeysTable)

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
	query := fmt.Sprintf("UPDATE %s "+
		"SET `banned` = 0 WHERE id = ?", licenseKeysTable)

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
	query := fmt.Sprintf("UPDATE %s "+
		"SET `is_deleted` = 1 WHERE id = ?", licenseKeysTable)

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
	query := fmt.Sprintf("UPDATE %s "+
		"SET `hwid` = null , hwidk = null WHERE id = ?", licenseKeysTable)

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
