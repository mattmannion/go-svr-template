package env

import (
	"fmt"

	"github.com/spf13/viper"
)

type Cfg struct {
	Env          string `mapstructure:"ENV"`
	Port         string `mapstructure:"PORT"`
	DSN          string `mapstructure:"DSN"`
	Session_Name string `mapstructure:"SESSION_NAME"`
	Redis_Addr   string `mapstructure:"REDIS_ADDR"`
	Redis_Secret string `mapstructure:"REDIS_SECRET"`
}

var CFG Cfg

func LoadConfig() Cfg {
	viper.SetConfigFile("./src/env/.env")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	err = viper.Unmarshal(&CFG)

	if err != nil {
		fmt.Println(err)
	}

	return CFG
}
