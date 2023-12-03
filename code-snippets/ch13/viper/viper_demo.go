package main

// viper demo

import (
	"fmt"
	"github.com/spf13/viper"
)


// envDemo 从环境变量读取配置
func envDemo(){
	viper.SetEnvPrefix("spf")
	viper.AutomaticEnv()
	//viper.BindEnv("http_port")
	fmt.Println(viper.Get("name"))
	fmt.Println(viper.AllKeys())
}

func main() {
	envDemo()
}
