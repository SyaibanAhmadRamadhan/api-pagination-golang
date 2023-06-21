package config

import (
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg    Config
	doOnce sync.Once
)

type Config struct {
	Application struct {
		Port string `mapstructure:"PORT"`
	} `mapstructure:"APPLICATION"`

	DB struct {
		Mysql struct {
			Host string `mapstructure:"HOST"`
			Port int    `mapstructure:"PORT"`
			User string `mapstructure:"USER"`
			Pass string `mapstructure:"PASS"`
			Name string `mapstructure:"NAME"`
		} `mapstructure:"MYSQL"`
	} `mapstructure:"DB"`
}

func Get() Config {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read file env %v", err)
	}

	doOnce.Do(func() {
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("cannot unmarshal : %v", err)
		}
	})

	return cfg
}
