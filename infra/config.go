package infra

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitConfig() {
	initViper()
}

func initViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read file config: %v", err))
	}

	for _, key := range viper.AllKeys() {
		value := viper.Get(key)
		viper.SetDefault(key, value)
	}
}
