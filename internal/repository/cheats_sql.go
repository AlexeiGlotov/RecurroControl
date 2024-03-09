package repository

import (
	"database/sql"
	"fmt"

	"RecurroControl/models"
)

type CheatSql struct {
	db *sql.DB
}

func NewCheatSql(db *sql.DB) *CheatSql {
	return &CheatSql{db: db}
}

func (a *CheatSql) GetCheats() ([]models.Cheats, error) {

	cheats := []models.Cheats{}

	query := fmt.Sprintf("SELECT * FROM %s", cheatTable)
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
