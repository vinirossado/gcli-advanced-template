package main

import (
	"basic/pkg/config"
	"basic/pkg/logger"

	"context"
	"flag"
)

func main() {
	//conf := config.NewConfig()
	//log := logger.NewLog(conf)
	//
	//dbType := repository.PostgreSQL
	//
	//_ = repository.NewDB(dbType, conf)
	//
	//app, cleanup, err := newApp(dbType, conf, log)
	//
	//if err != nil {
	//	panic(err)
	//}
	//app.Run()
	//
	////if it needs to drop all table from DB
	////app.DropAll()
	//
	//defer cleanup()

	var envConf = flag.String("conf", "../../../config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
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
