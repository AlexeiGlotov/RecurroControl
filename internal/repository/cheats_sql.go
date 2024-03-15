package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"RecurroControl/models"
)

type CheatSql struct {
	db *sql.DB
}

func NewCheatSql(db *sql.DB) *CheatSql {
	return &CheatSql{db: db}
}

func (a *CheatSql) CreateCheat(cheat *models.Cheats) (int, error) {

	query := fmt.Sprintf("INSERT INTO %s (name, secure, is_allowed_generate) VALUES (?, ?, ?)", cheatTable)
	result, err := a.db.Exec(query, cheat.Name, cheat.Secure, cheat.IsAllowedGenerate)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (a *CheatSql) UpdateCheat(cheat *models.Cheats) error {

	if cheat.Id == 0 {
		return errors.New("id cheat 0")
	}
	query := fmt.Sprintf("UPDATE %s SET name = ?, secure = ?, is_allowed_generate = ? WHERE id = ?", cheatTable)
	result, err := a.db.Exec(query, cheat.Name, cheat.Secure, cheat.IsAllowedGenerate, cheat.Id)
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

func (a *CheatSql) GetCheats(role string) ([]models.Cheats, error) {

	cheats := []models.Cheats{}
	var query string

	switch role {
	case models.Admin:
		query = fmt.Sprintf("SELECT * FROM %s", cheatTable)
	case models.Distributors, models.Reseller:
		query = fmt.Sprintf("SELECT * FROM %s WHERE is_allowed_generate = 1", cheatTable)
	default:
		return nil, errors.New("bad role")
	}

	row, err := a.db.Query(query)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := models.Cheats{}
		err := row.Scan(&temp.Id, &temp.Name, &temp.Secure, &temp.IsAllowedGenerate)
		if err != nil {
			return nil, err
		}
		cheats = append(cheats, temp)
	}

	return cheats, nil
}
