package configuration

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var Configuration AppConfig

type Application struct {
	Address string `mapstructure:"address"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

type AppConfig struct {
	MicroserviceName    string      `mapstructure:"microserviceName"`
	MicroserviceServer  string      `mapstructure:"microserviceServer"`
	MicroservicePort    string      `mapstructure:"microservicePort"`
	MicroserviceVersion string      `mapstructure:"microserviceVersion"`
	Application         Application `mapstructure:"application"`
	Database            Database    `mapstructure:"database"`
	LogLevel            string      `mapstructure:"logLevel"`
}

func LoadAppConfiguration() {
	fmt.Println("Loading configuration from file [app.yml]")

	// Set the file name of the configurations file
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		panic("Error reading config file")
	}

	// Set undefined variables
	hostname, _ := os.Hostname()
	viper.SetDefault("microserviceServer", hostname)

	err := viper.Unmarshal(&Configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		panic(fmt.Sprintf("Unable to decode into struct, %v", err))
	}

	fmt.Println(Configuration)

	fmt.Println("Configuration loaded")
}
