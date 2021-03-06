package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type EnvConfig struct {
	Port     string `mapstructure:"PORT"`
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBName   string `mapstructure:"DB_NAME"`
}

var ENV *EnvConfig

func Load() {
	ENV = &EnvConfig{}
	env := strings.ToLower(os.Getenv("GO_ENVIRONMENT"))
	viper.SetConfigFile(fmt.Sprintf("%s.env", env))
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var viperErr viper.ConfigFileNotFoundError
		if ok := errors.As(err, &viperErr); ok {
			log.Fatalln(fmt.Errorf("config file not found. %w", err))
		} else {
			log.Fatalln(fmt.Errorf("unexpected error loading config file. %w", err))
		}
	}

	if err := viper.Unmarshal(ENV); err != nil {
		log.Fatalln(fmt.Errorf("failed to unmarshal config. %w", err))
	}

	fmt.Println("Success: environment config file loaded.")
}
