package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// MainConfig is overall configuration structure
type MainConfig struct {
	Duration   int    `mapstructure:"duration"`
	ModMessage string `mapstructure:"modmessage"`
	BOT_TOKEN  string
}

// C is the mainconfig object
var C MainConfig

func init() {
	viper.AutomaticEnv()
	C.BOT_TOKEN = viper.GetString("BOT_TOKEN")
}

// ReadConfig sets configuration of complete environment
func ReadConfig() MainConfig {

	viper := viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return C
}
