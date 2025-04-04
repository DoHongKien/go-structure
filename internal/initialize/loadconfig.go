package initialize

import (
	"fmt"

	"github.com/DoHongKien/go-structure/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()

	viper.AddConfigPath("./config")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unabel to decode configuration %v", err)
	}
}
