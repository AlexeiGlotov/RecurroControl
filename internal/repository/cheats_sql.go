package repository

import (
	"database/sql"
	"fmt"

	todo "RecurroControl"
)

type CheatSql struct {
	db *sql.DB
}

func NewCheatSql(db *sql.DB) *CheatSql {
	return &CheatSql{db: db}
}

func (a *CheatSql) GetCheats() ([]todo.StCheats, error) {

	cheats := []todo.StCheats{}

	query := fmt.Sprintf("SELECT * FROM %s", cheatTable)
	row, err := a.db.Query(query)

	if err != nil {
		return nil, err
	}

	for row.Next() {
		temp := todo.StCheats{}
		err := row.Scan(&temp.Id, &temp.Name, &temp.Secure, &temp.IsAllowedGenerate)
		if err != nil {
			return nil, err
		}
		cheats = append(cheats, temp)
	}

	return cheats, nil
}
