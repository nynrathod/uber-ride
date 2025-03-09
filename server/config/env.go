package config

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
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
}

// Call to load the variables from env
func loadEnvVariables() (config *envConfigs) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
