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
	Port           string `mapstructure:"PORT"`
	DatabaseDriver string `mapstructure:"DB_DRIVER"`
	DatabaseName   string `mapstructure:"DB_NAME"`
}

var ENV *EnvConfig

func Load() {
	ENV := &EnvConfig{}
	environment := strings.ToLower(os.Getenv("GO_ENVIRONMENT"))

	viper.SetConfigFile(fmt.Sprintf("%s.env", environment))
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var viperErr viper.ConfigFileNotFoundError
		if ok := errors.As(err, &viperErr); !ok {
			log.Fatal(fmt.Errorf("config file not found %s", err))
		} else {
			log.Fatal(fmt.Errorf("unexpected error %w", err))
		}
	}

	if err := viper.Unmarshal(ENV); err != nil {
		log.Fatal(fmt.Errorf("failed to unmarshal config %w", err))
	}

	fmt.Println("Success: environment config loaded.")
}
