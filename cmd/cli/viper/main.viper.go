package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	}
	Mysql struct {
		Host            string `mapstructure:"host"`
		Port            int    `mapstructure:"port"`
		Username        string `mapstructure:"username"`
		Password        string `mapstructure:"password"`
		Dbname          string `mapstructure:"dbname"`
		MaxIdleConns    int    `mapstructure:"maxIdleConns"`
		MaxOpenConns    int    `mapstructure:"maxOpenConns"`
		ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
	}
	Logger struct {
		LogLevel    string `mapstructure:"debug"`
		FileLogName string `mapstructure:"file_log_name"`
		MaxBackups  int    `mapstructure:"max_backups"`
		MaxAge      int    `mapstructure:"max_age"`
		MaxSize     int    `mapstructure:"max_size"`
		Compress    bool   `mapstructure:"compress"`
	}
}

func main() {

	viper := viper.New()

	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println("Server Port: ", viper.GetInt("server.port"))

	// config structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	fmt.Printf("config database info: %+v\n", config.Mysql)
}
