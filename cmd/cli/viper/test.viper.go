package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// mapstructure decoder (what viper.Unmarshal uses) which key in the config maps to that field.
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()

	viper.AddConfigPath("./config/") // add path to config
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration: %w \n", err))
	}
	// read server configuration
	fmt.Println("Server Port::", viper.GetInt("server.port"))

	// read config into object

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to decode configuration %v \n", err)
	}

	fmt.Println(config)
}
