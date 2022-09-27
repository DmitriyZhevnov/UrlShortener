package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Domain  string
	Storage Storage
	HTTP    HTTP
}

type HTTP struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
}

type Storage struct {
	Postgresql Postgresql
	Redis      Redis
}

type Redis struct {
	Addr     string
	Port     string
	Password string
	DB       int
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

	cfg.Domain = viper.GetString("DOMAIN")

	cfg.HTTP.Port = viper.GetString("PORT")
	cfg.HTTP.ReadTimeout = viper.GetInt("READ_TIMEOUT")
	cfg.HTTP.WriteTimeout = viper.GetInt("WRITE_TIMEOUT")

	cfg.Storage.Postgresql.Host = viper.GetString("POS_HOST")
	cfg.Storage.Postgresql.Port = viper.GetString("POS_PORT")
	cfg.Storage.Postgresql.Database = viper.GetString("POS_DATABASE")
	cfg.Storage.Postgresql.Username = viper.GetString("POS_USERNAME")
	cfg.Storage.Postgresql.Password = viper.GetString("POS_PASSWORD")

	cfg.Storage.Redis.Addr = viper.GetString("REDIS_IMAGE")
	cfg.Storage.Redis.Port = viper.GetString("REDIS_PORT")
	cfg.Storage.Redis.Password = viper.GetString("REDIS_PASSWORD")
	cfg.Storage.Redis.DB = viper.GetInt("REDIS_DB")

	return nil
}
