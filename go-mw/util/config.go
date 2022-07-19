package util

import "github.com/spf13/viper"

type GlobalConfig struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	TargetAddress string `mapstructure:"TARGET_ADDRESS""`
}

func LoadGlobalConfig(path string) (config GlobalConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
