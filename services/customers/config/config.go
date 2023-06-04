package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerSettings
	Database Database
}

type ServerSettings struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func setupConfigFile() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic("The config file was not found")
	}
	return nil
}

func getConfigParameter(parameter string) string {
	return viper.GetString(parameter)
}

func GetConfig() Config {
	setupConfigFile()
	return Config{
		Server: ServerSettings{
			Port: getConfigParameter("server.port"),
		},
		Database: Database{
			Host:     getConfigParameter("db.host"),
			Port:     getConfigParameter("db.port"),
			Database: getConfigParameter("db.database"),
			User:     getConfigParameter("db.user"),
			Password: getConfigParameter("db.password"),
		},
	}
}
