package internal

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Temporal struct {
		Port            string
		Host            string
		reconnect_delay int
	}
}

var configHolder *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigFile("config.yaml")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Failed to read config file: %s", err)
		}

		config := &Config{}
		err = viper.Unmarshal(config)
		if err != nil {
			log.Fatalf("Failed to unmarshal configuration: %s", err)
		}

		configHolder = config
	})

	log.Default().Print("Fetched config")
	return configHolder
}
