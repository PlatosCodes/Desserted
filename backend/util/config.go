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
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Try reading the configuration from the file
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore this error and continue
			log.Printf("Config file not found, using environment variables")
		} else {
			// Some other error occurred while reading the config file
			log.Printf("Error reading config: %v", err)
			return
		}
	}
	// Unmarshal the configuration (from file and/or env vars) into the Config struct
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Printf("Error unmarshaling config: %v", err)
	}

	return
}
