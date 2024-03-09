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

func (a *CheatSql) CreateCheats(cheat *models.Cheats) (int, error) {

	var id int

	query := fmt.Sprintf("INSERT INTO %s (name ,secure,is_allowed_generate) values (?,?,?)", cheatTable)
	row := a.db.QueryRow(query, cheat.Name, cheat.Secure, cheat.IsAllowedGenerate)
	if row.Err() != nil {
		return 0, row.Err()
	}

	row = a.db.QueryRow(fmt.Sprintf("SELECT MAX(id) FROM %s;", cheatTable))

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *CheatSql) GetCheats(role string) ([]models.Cheats, error) {

	cheats := []models.Cheats{}
	var query string

	switch role {
	case models.Admin:
		query = fmt.Sprintf("SELECT * FROM %s", cheatTable)
	case models.Distributors:
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
