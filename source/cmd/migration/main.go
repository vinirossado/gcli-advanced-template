package main

import (
	"basic/pkg/config"
	logger "basic/pkg/logger"
)

func main() {
	conf := config.NewConfig()
	log := logger.NewLog(conf)

	app, cleanup, err := newApp(conf, log)
	if err != nil {
		panic(err)
	}
	app.Run()

	//if it needs to drop all table from DB
	//app.DropAll()

	defer cleanup()
}
