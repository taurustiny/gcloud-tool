package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	// "gopkg.in/yaml.v2"
)

func GetConfig() {
	viper.SetDefault("Network", "default")
	viper.SetDefault("SubNetwork", "default")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)         
	}
	// fmt.Println(viper.AllSettings())
	fmt.Println(viper.Get("gke").([]interface{})[0])

}
