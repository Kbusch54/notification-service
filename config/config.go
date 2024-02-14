package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	devPrefix     = "development"
	envProfileKey = "ENVIRONMENT"
)

var cfg *Config

func Load(cfgpath string) *Config {

	fmt.Println("Starting to load configuration")

	if cfg != nil {
		fmt.Println("Configuration already loaded")
		return cfg
	}

	viper.AddConfigPath(cfgpath)
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	fmt.Println("Attempting to read configuration file")
	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading configuration file: " + err.Error())
	}

	bindEnvVars()
	mergeProfileVars()
	if err := viper.Unmarshal(&cfg); err != nil {
		panic("Error unmarshalling configuration: " + err.Error())
	}
	return cfg
}

func mergeProfileVars() {
	env := os.Getenv(envProfileKey)

	switch {
	case env != "":
		viper.SetConfigName("app-" + env)
	default:
		viper.SetConfigName("app-" + devPrefix)
	}

	if err := viper.MergeInConfig(); err != nil {
		panic("Error merging profile config vars. Reason: " + err.Error())
	}
}

// bindEnvVars bind any entry env var
func bindEnvVars() {
	bindEnvVar("persistence.mongodb.url", "MONGODB_URI")
	bindEnvVar("services.brevo.apiKey", "BREVO_API_KEY")
	bindEnvVar("services.brevo.email", "BREVO_EMAIL")
}

func bindEnvVar(key string, envKey string) {
	err := viper.BindEnv(key, envKey)
	if err != nil {
		panic("Error binding env vars. Reason: " + err.Error())
	}
}
