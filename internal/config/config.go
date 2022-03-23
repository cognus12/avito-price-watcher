package config

import (
	"apricescrapper/pkg/logger"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host   string `env:"HOST" env-default:"localhost"`
	Port   string `env:"PORT" env-default:"3000"`
	DbPath string `env:"DB_PATH"`
}

var instance *Config
var once sync.Once

func GetConfig(logger logger.Logger) *Config {
	once.Do(func() {

		instance = &Config{}

		if err := cleanenv.ReadConfig(".env", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err.Error())
		}
	})
	return instance
}
