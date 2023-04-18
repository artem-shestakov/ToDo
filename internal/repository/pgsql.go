package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Address  string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPgsqlDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.Address, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
		),
	)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Can't ping database: %s", err)
		return nil, err
	}
	return db, nil
}
