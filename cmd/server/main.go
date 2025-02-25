package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	initViperWatch()
	app := InitWebServer()

	err := app.Server.Run(":8080")
	if err != nil {
		panic(err)
	}

	app.Server.Run(":8080")

}

func initViperWatch() {
	cfile := pflag.String("config", "./config/dev.yaml", "配置文件路径")
	pflag.Parse()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(*cfile)
	viper.WatchConfig()

	//读取配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
