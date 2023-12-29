package main

import (
	"basic/pkg/cache"
	"basic/pkg/config"
	"basic/pkg/http"
	"basic/pkg/logger"
	"basic/source/repository"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	log := logger.NewLog(conf)
	cache.MemoryCache()
	dbType := repository.PostgreSQL

	_ = repository.NewDB(dbType, conf)

	app, cleanup, err := newApp(dbType, conf, log)

	if err != nil {
		panic(err)
	}
	log.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
	defer cleanup()
}
