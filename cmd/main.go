package main

import (
	"fmt"

	"github.com/artem-shestakov/to-do/internal/config"
	"github.com/artem-shestakov/to-do/internal/handler"
	"github.com/artem-shestakov/to-do/internal/repository"
	"github.com/artem-shestakov/to-do/internal/server"
	"github.com/artem-shestakov/to-do/internal/service"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	logger.SetFormatter(&logrus.JSONFormatter{})

	config, err := config.ReadConfig("./config.yaml", logger)
	if err != nil {
		logger.Fatalf("Config read error")
	}

	db, err := repository.NewPgsqlDB(repository.Config{
		Address:  config.Database.Address,
		Port:     config.Database.Port,
		User:     config.Database.Username,
		Password: config.Database.Password,
		DBName:   config.Database.DBName,
		SSLMode:  "disable",
	})
	if err != nil {
		logger.Fatalf("Can't ping database: %s", err)
	}

	repository.RunDBMigration(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.Database.Username,
			config.Database.Password,
			config.Database.Address,
			config.Database.Port,
			config.Database.DBName,
		),
		logger,
	)

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHadler(service, logger, config.APIToken)
	server := server.NewServer(logger)
	if err := server.Run(config.Server.Address, config.Server.Port, handler.InitRouters()); err != nil {
		logger.Fatalf("Can't run server. Got error: %s", err.Error())
	}
}
