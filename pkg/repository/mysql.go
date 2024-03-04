package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string
	Port     string
	Login    string
	Password string
	DBName   string
	SSLMode  string
}

func NewMysqlDB(cfg Config) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Login,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
