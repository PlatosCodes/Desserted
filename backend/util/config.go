package util

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// TODO: Look into changing TokenSymmetricKey to Asymmetric in production.
type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	TestDBSource         string        `mapstructure:"TEST_DB_SOURCE"`
	MigrationURL         string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	FrontendAddress      string        `mapstructure:"FRONTEND_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	SmtpHost             string        `mapstructure:"smtpHost"`
	SmtpPort             int           `mapstructure:"smtpPort"`
	SmtpUsername         string        `mapstructure:"smtpUsername"`
	SmtpPassword         string        `mapstructure:"smtpPassword"`
	SmtpSender           string        `mapstructure:"smtpSender"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AutomaticEnv() // Read from environment variables first

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig() // This is now the fallback option
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if we have all needed env vars
			err = nil
		} else {
			// Config file was found but another error was produced
			return
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("error unmarshaling config: %v", err)
	}
	return
}
