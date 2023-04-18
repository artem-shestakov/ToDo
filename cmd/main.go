package main

import (
	"github.com/artem-shestakov/to-do/internal/handler"
	"github.com/artem-shestakov/to-do/internal/repository"
	"github.com/artem-shestakov/to-do/internal/server"
	"github.com/artem-shestakov/to-do/internal/service"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	logger.SetFormatter(&logrus.JSONFormatter{})
	// log.SetFormatter(&log.JSONFormatter{})

	db, err := repository.NewPgsqlDB(repository.Config{
		Address:  "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		logger.Fatalf("Can't ping database: %s", err)
	}
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHadler(service, logger)
	server := server.NewServer(logger)
	if err := server.Run("0.0.0.0", "8000", handler.InitRouters()); err != nil {
		logger.Fatalf("Can't run server. Got error: %s", err.Error())
	}
}
