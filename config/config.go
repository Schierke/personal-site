package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type SourceCfg struct {
	SourcePath string `mapstructure:"SOURCE_PATH"`
}

type AppConfig struct {
	Source SourceCfg `mapstructure:"source"`
}

func LoadAppConfig(path string) (AppConfig, error) {
	if path == "" {
		return AppConfig{}, errors.New("config path is empty")
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return AppConfig{}, errors.New("config file not found")
		}

		return AppConfig{}, fmt.Errorf("failed to read config file: %s", err.Error())
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return AppConfig{}, fmt.Errorf("failed to unmarshal config: %s", err.Error())
	}

	return config, nil
}
