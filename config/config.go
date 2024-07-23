package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"project-management-service/pkg/Logging"
)

type Config struct {
	Server Server `yaml:"server"`
	DB     DB     `yaml:"db"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

var instance *Config

func GetConfig() *Config {
	logger := Logging.GetLogger()
	logger.Info("read application configuration")
	instance = &Config{}
	if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)
		logger.Info(help)
		logger.Fatal(err)
	}

	return instance
}
