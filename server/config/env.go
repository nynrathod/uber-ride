package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

// InitEnvConfigs loads environment variables and returns the configuration.
func InitEnvConfigs() *envConfigs {
	EnvConfigs = loadEnvVariables()
	return EnvConfigs
}

type envConfigs struct {
	DbName          string `mapstructure:"DB_NAME"`
	DbHost          string `mapstructure:"DB_HOST"`
	DbUser          string `mapstructure:"DB_USER"`
	DbPassword      string `mapstructure:"DB_PASSWORD"`
	DbPort          string `mapstructure:"DB_PORT"`
	JWTSecrete      string `mapstructure:"JWT_SECRET"`
	SmtpEmail       string `mapstructure:"SMTP_EMAIL"`
	SmtpToken       string `mapstructure:"SMTP_TOKEN"`
	SmtpToEmail     string `mapstructure:"SMTP_TO_EMAIL"`
	VerrifyOtpToken string `mapstructure:"VERIFY_OTP_TOKEN"`
	ENV             string `mapstructure:"ENV"`
	HTTPS           string `mapstructure:"HTTPS"`
}

// loadEnvVariables loads the environment variables from .env and unmarshals them.
func loadEnvVariables() *envConfigs {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Viper reads the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Unmarshal the environment variables into the struct.
	var config envConfigs
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}
