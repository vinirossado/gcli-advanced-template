package main

import (
	"basic/pkg/config"
	"basic/pkg/http"
	"basic/pkg/logger"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	log := logger.NewLog(conf)

	log.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	app, cleanup, err := newApp(conf, log)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
}
