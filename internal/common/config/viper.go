package config

import (
	"github.com/spf13/viper"
)

func NewViperConfig() error {

	viper.SetConfigName("global")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../common/config")  // 在 kitchen 目录下调用, 相对路径这么写合理
	viper.AutomaticEnv()
	
	return viper.ReadInConfig()
}