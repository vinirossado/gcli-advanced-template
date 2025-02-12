package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"

	"basic/pkg/config"
	"basic/pkg/logger"
)

func main() {
	var envConf = flag.String("conf", "", "config path, eg: -conf ./config/local.yml")
	flag.Parse()

	if *envConf == "" {
		// Determine the correct path based on the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		if filepath.Base(cwd) == "server" {
			*envConf = "../../../config/local.yml"
		} else {
			*envConf = "./config/local.yml"
		}
	}

	fmt.Printf("Using config file: %s\n", *envConf)

	conf := config.NewConfig(*envConf)

	log := logger.NewLog(conf)

	app, cleanup, err := NewWire(conf, log)
	defer cleanup()
	if err != nil {
		log.Fatal("failed to initialize application", zap.Error(err))
	}

	log.Info("server start", zap.String("host", fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))))
	log.Info("docs addr", zap.String("addr", fmt.Sprintf("http://%s:%d/swagger/index.html", conf.GetString("http.host"), conf.GetInt("http.port"))))

	if err = app.Run(context.Background()); err != nil {
		log.Fatal("application run failed", zap.Error(err))
	}
}
