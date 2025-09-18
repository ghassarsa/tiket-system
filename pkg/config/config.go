package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`

	ServerAddress string `mapstructure:"SERVER_ADDRESS"`

	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`

	JWTSecret string `mapstructure:"JWT_SECRET"`

	DefaultDB *gorm.DB
}

func LoadConfig() (Config, error) {
	var cfg Config

	v := viper.NewWithOptions(viper.ExperimentalBindStruct())
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
