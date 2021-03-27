package config

import "github.com/spf13/viper"

// AppConfig stores all configuration of the application.
type AppConfig struct {
	Port     int    `mapstructure:"PORT"`
	Host     string `mapstructure:"HOST"`
	DBSource string `mapstructure:"DB_SOURCE"`
	DBName   string `mapstructure:"DB_NAME"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string, config *AppConfig) error {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}

	return nil
}
