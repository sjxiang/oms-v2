package conf

import (
	"github.com/spf13/viper"
)

func NewViperConfig() error {

	viper.SetConfigName("global")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../common/conf")  // 在 order 目录下调用, 相对路径这么写合理
	viper.AutomaticEnv()
	
	return viper.ReadInConfig()
}