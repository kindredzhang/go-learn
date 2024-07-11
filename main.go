package main

import (
	"fmt"
	"github.com/spf13/viper"
	"go-learn/config"
)

func main() {
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var c config.Server
	err := viper.Unmarshal(&c)
	if err != nil {
		return
	}
	fmt.Println(c.Mysql)
}
