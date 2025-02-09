package config

import (
	"github.com/spf13/viper"
)

type conf struct {
	Workers        int    `mapstructure:"workers"`
	ScrapeInterval int    `mapstructure:"scrape_interval_sec"`
	BackendURL     string `mapstructure:"backend_url"`
}

var C conf

func MustInit() {
	viper.SetDefault("workers", 5)
	viper.BindEnv("workers", "WORKERS")

	viper.SetDefault("scrape_interval_sec", 10)
	viper.BindEnv("scrape_interval_sec", "SCRAPE_INTERVAL_SEC")

	viper.SetDefault("backend_url", "http://localhost:8082")
	viper.BindEnv("backend_url", "BACKEND_URL")

	viper.AutomaticEnv()

	if err := viper.UnmarshalExact(&C); err != nil {
		panic(err)
	}
}
