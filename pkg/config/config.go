package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

func NewConfig(p string) *viper.Viper {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		envConf = p
	}
	if envConf == "" {
		envConf = "../config/local.yml"
	}

	basepath := getConfigPath() + "/config/local.yml"

	return getConfig(basepath)
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}

func getConfigPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	if runtime.GOOS == "windows" {
		basepath = strings.Replace(basepath, "\\", "/", -1)
	}
	index := strings.LastIndex(basepath, "/pkg")
	if index != -1 {
		basepath = basepath[:index]
	}
	return basepath
}
