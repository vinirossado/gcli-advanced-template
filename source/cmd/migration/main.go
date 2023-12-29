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
	gormDb, err := repository.NewDB(dbType, conf)

	db, err := gormDb.Connect()

	if err != nil {
		panic(err)
	}

	_ = repository.NewRepository(log, db)

	app, cleanup, err := newApp(dbType, conf, db, log)

	if err != nil {
		panic(err)
	}
	app.Run()

	//if it needs to drop all table from DB
	//app.DropAll()

	defer cleanup()
}
