package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func NewConfig(p string) *viper.Viper {
	// APP_CONF env var takes priority over the flag/default path
	if envConf := os.Getenv("APP_CONF"); envConf != "" {
		p = envConf
	}
	return getConfig(p)
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	if err := conf.ReadInConfig(); err != nil {
		log.Fatalf("failed to load config file %q: %v\n\nSet APP_CONF to the absolute path of your config file.", path, err)
	}
	fmt.Printf("config loaded: %s\n", path)
	return conf
}
