package config


import "github.com/spf13/viper"

	
/*

	- internal

		- stock
			- main.go

		- common
			- conf
				- conf.yaml

	略

 */


const (
	path = "../common/config"
)

func NewViperConfig() error {

	viper.AddConfigPath(path)       // 配置文件所在目录
	viper.SetConfigName("config")   // 文件名（不带扩展名）
	viper.SetConfigType("yaml")     // 显式指定格式

	return viper.ReadInConfig()     // 读取文件内容
}