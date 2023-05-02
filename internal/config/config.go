package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Address string `yaml:"address" envconfig:"SERVER_ADDRESS"`
		Port    string `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	Database struct {
		Address  string `yaml:"address" envconfig:"DB_ADDRESS"`
		Port     string `yaml:"port" envconfig:"DB_PORT"`
		Username string `yaml:"username" envconfig:"DB_USER"`
		Password string `yaml:"password" envconfig:"DB_PASS"`
		DBName   string `yaml:"db_name" envconfig:"DB_NAME"`
	} `yaml:"database"`
	APIToken string `yaml:"api_token" envconfig:"API_TOKEN"`
}

func ReadConfig(path string, logger *logrus.Logger) (*Config, error) {
	f, err := os.Open(path)
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

func ReadEnv(cfg *Config) error {
	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}
	return nil
}
