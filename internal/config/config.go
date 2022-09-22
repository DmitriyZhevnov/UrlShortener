package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Storage Storage
}

type Storage struct {
	Postgresql Postgresql
}

type Postgresql struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := fromEnv(instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}

func fromEnv(cfg *Config) error {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	cfg.Storage.Postgresql.Host = viper.GetString("HOST")
	cfg.Storage.Postgresql.Port = viper.GetString("PORT")
	cfg.Storage.Postgresql.Database = viper.GetString("DATABASE")
	cfg.Storage.Postgresql.Username = viper.GetString("USERNAME")
	cfg.Storage.Postgresql.Password = viper.GetString("PASSWORD")

	return nil
}
