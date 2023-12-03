package env

import (
	"log"

	"github.com/spf13/viper"
)

type ENV struct {
	AppPort       string `mapstructure:"APP_PORT"`
	AppVersion    string `mapstructure:"APP_VERSION"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_USER"`
	RedisUsername string `mapstructure:"REDIS_PASS"`
}

func NewEnv() *ENV {
	env := ENV{}
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment can't be loaded: ", err)
	}

	return &env
}
