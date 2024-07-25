package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
        Port int `mapstruct:"port"`
    } `mapstruct:"server"`

	Databases []struct {
		User string `mapstruct:"user"`
		Password string `mapstruct:"password"`
		Host string `mapstruct:"host"`
		DbName string `mapstruct:"dbName"`
	} `mapstruct:"databases"`
}

func main() {
	viper := viper.New()
	
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// Read config file
	err := viper.ReadInConfig()

	if err != nil {
        panic(fmt.Errorf("Error reading config file: " + err.Error()))
    }

	// Read server config
	fmt.Println("Server port: ", viper.GetInt("server.port"))
	fmt.Println("Security: ", viper.GetString("security.jwt.key"))

	// Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config: %v", err)
	}

	fmt.Println("Server port: ", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("Database: %s:%s@%s/%s\n", db.User, db.Password, db.Host, db.DbName)
	}
}