package shared

import (
	"github.com/spf13/viper"
)

type Config struct {
	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`
	DBSource  string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// or using -> viper.SetConfigFile(".env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// overide the variable in env instead of from the config file
	// viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}
