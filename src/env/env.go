package env

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
	DSN  string `mapstructure:"DSN"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
