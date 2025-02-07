package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type conf struct {
	Port        int    `mapstructure:"port"`
	PostgresURL string `mapstructure:"postgres_url"`
}

var C conf

func MustInit() {
	viper.SetDefault("port", 8082)

	viper.BindEnv("port", "PORT")
	viper.BindEnv("postgres_url", "POSTGRES_URL")

	viper.AutomaticEnv()

	if err := viper.UnmarshalExact(&C); err != nil {
		panic(err)
	}

	fmt.Println(C)
}
