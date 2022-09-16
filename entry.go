package main

import (
	"log"
	"xdsec-join/src/cache"
	"xdsec-join/src/config"
	"xdsec-join/src/controller"
	logger2 "xdsec-join/src/logger"
	"xdsec-join/src/repository"
)

func main() {
	if err := config.Initialize(); err != nil {
		log.Panic("init config error: ", err)
	}
	if err := logger2.Initialize(); err != nil {
		log.Panic("init logger error: ", err)
	}
	logger2.Info("server is initializing...")
	if err := cache.Initialize(); err != nil {
		logger2.PanicAny("init cache error: ", err)
	}
	if err := repository.Initialize(); err != nil {
		logger2.PanicAny("init repository error: ", err)
	}
	if err := controller.Initialize(); err != nil {
		logger2.PanicAny("init controller error: ", err)
	}
	logger2.Info("initialization completed, starting service...")
	if err := controller.Run(); err != nil {
		logger2.PanicAny("run server error: ", err)
	}
}
