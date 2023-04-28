package config

import (
	"os"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Address string `yaml:"address"`
		Port    string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Address  string `yaml:"address"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db_name"`
	} `yaml:"database"`
	APIToken string `yaml:"api_token"`
}

func ReadConfig(path string, logger *logrus.Logger) (*Config, error) {
	f, err := os.Open("config.yml")
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Errorf("Can't read config file")
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Errorf("Can't parse config file")
		return nil, err
	}
	return &cfg, nil
}
