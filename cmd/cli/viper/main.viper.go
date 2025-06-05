package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	}
	Databases []struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Dbname   string `mapstructure:"dbname"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	var config Config
	// đọc config
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read in YAML file: %w", err))
	}

	// fmt.Println("Server Port: ", viper.GetInt("server.port"))
	// fmt.Println("Database Name: ", viper.GetString("databases.mysql.dbname"))

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("failed to unmarshal YAML file: %w", err))
	}

	fmt.Println("Server Port: ", config.Server.Port)

	for _, database := range config.Databases {
		// fmt.Println("Database Username: ", database.Username, "Database Name: ", database.Dbname, "Database Port: ", database.Port, "Database Host: ", database.Host)
		fmt.Printf("Database Username: %s, Name: %s, Port: %d, Host: %s\n", database.Username, database.Dbname, database.Port, database.Host)
	}
}
