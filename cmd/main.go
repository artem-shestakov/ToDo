package main

import (
	"flag"
	"fmt"

	"github.com/artem-shestakov/to-do/internal/config"
	"github.com/artem-shestakov/to-do/internal/handler"
	"github.com/artem-shestakov/to-do/internal/repository"
	"github.com/artem-shestakov/to-do/internal/server"
	"github.com/artem-shestakov/to-do/internal/service"
	"github.com/sirupsen/logrus"
)

var (
	confPath *string
	logger   = logrus.New()
)

func init() {
	confPath = flag.String("config", "config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()
	logger.SetFormatter(&logrus.JSONFormatter{})

	conf, err := config.ReadConfig(*confPath, logger)
	if err != nil {
		logger.Fatalf("Config read error")
	}
	if err = config.ReadEnv(conf); err != nil {
		logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Fatalf("Can't parse config file")
	}

	db, err := repository.NewPgsqlDB(repository.Config{
		Address:  conf.Database.Address,
		Port:     conf.Database.Port,
		User:     conf.Database.Username,
		Password: conf.Database.Password,
		DBName:   conf.Database.DBName,
		SSLMode:  "disable",
	})
	if err != nil {
		logger.Fatalf("Can't ping database: %s", err)
	}

	repository.RunDBMigration(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			conf.Database.Username,
			conf.Database.Password,
			conf.Database.Address,
			conf.Database.Port,
			conf.Database.DBName,
		),
		logger,
	)

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHadler(service, logger, conf.APIToken)
	server := server.NewServer(logger)
	if err := server.Run(conf.Server.Address, conf.Server.Port, handler.InitRouters()); err != nil {
		logger.Fatalf("Can't run server. Got error: %s", err.Error())
	}
}
