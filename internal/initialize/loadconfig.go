package initialize

import (
	"fmt"

	"github.com/longtk26/go-ecommerce/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
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

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Error unmarshalling config: %v", err)
	}
}
