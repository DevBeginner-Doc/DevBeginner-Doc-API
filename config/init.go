package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("[Config] 读取配置文件失败！")
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
