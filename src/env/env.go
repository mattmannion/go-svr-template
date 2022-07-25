package env

import (
	"github.com/spf13/viper"
)

type Cfg struct {
	Env  string `mapstructure:"ENV"`
	Port string `mapstructure:"PORT"`
	DSN  string `mapstructure:"DSN"`
}

func LoadConfig() (c Cfg, err error) {
	viper.SetConfigFile("./src/env/.env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	if err != nil {
		return
	}

	return
}
