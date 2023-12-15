package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Port                string `mapstructure:"PORT"`
	EmailSenderName     string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress  string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	DatabaseUrl         string `mapstructure:"DATABASE_URL"`
}

func LoadConfig() (config Config, err error) {

	viper.SetConfigFile("././.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&config)
	return
}
