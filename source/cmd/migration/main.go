package main

import (
	"basic/pkg/config"
	logger "basic/pkg/logger"
	"basic/source/repository"
)

func main() {
	conf := config.NewConfig()
	log := logger.NewLog(conf)

	dbType := repository.PostgreSQL
	
	_ = repository.NewDB(dbType, conf)

	app, cleanup, err := newApp(dbType, conf, log)

	if err != nil {
		panic(err)
	}
	app.Run()

	//if it needs to drop all table from DB
	//app.DropAll()

	defer cleanup()
}
