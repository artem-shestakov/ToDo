package repository

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
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

func RunDBMigration(sourceURL, dbURL string, logger *logrus.Logger) {
	m, err := migrate.New(sourceURL, dbURL)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatalf("Can't create migrate instance")
	}
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			logger.Info("DB migrations: No change")
			return
		}
		logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatalf("Can't migrate")
	}
	logger.Info("DB migrated succssfully")
}
