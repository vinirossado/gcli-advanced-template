package main

import (
	"basic/pkg/config"
	"basic/pkg/logger"
	"fmt"
	"os"
	"path/filepath"

	"context"
	"flag"
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
		panic(err)
	}
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
