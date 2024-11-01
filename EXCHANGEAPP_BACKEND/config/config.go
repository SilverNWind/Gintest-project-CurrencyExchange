package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Redis struct {
		Addr     string
		DB       int
		Password string
	}
}

var Appconfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		/*%v为相应值的默认格式*/
		log.Fatalf("Error reading config file: %v", err)
	}

	Appconfig = &Config{}

	if err := viper.Unmarshal(Appconfig); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	InitDB()
	InitRedis()
}
