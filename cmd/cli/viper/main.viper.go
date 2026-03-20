package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("local") // ten file config
	viper.SetConfigType("yaml")  // loai file config
	//read config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	//read server configuration
	fmt.Println("Server Port:", viper.GetInt("server.port"))
	fmt.Println("Server Security key:", viper.GetString("security.jwt.key"))
}
