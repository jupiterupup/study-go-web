package viperhelper

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetLocalConfIfPresent(key string) string {
	return viper.GetString(key)
}
