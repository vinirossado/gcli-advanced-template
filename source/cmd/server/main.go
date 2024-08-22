package main

import (
	"basic/pkg/config"
	"basic/pkg/logger"
	"context"
	"flag"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	log := logger.NewLog(conf)

	app, cleanup, err := NewWire(conf, log)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	log.Info("server start", zap.String("host", fmt.Sprintf("http://%s:%d", conf.GetString("http.host"), conf.GetInt("http.port"))))
	log.Info("docs addr", zap.String("addr", fmt.Sprintf("http://%s:%d/swagger/index.html", conf.GetString("http.host"), conf.GetInt("http.port"))))
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
