package main

import (
	"XDSEC2022-Backend/src/cache"
	"XDSEC2022-Backend/src/config"
	"XDSEC2022-Backend/src/controller"
	"XDSEC2022-Backend/src/logger"
	"XDSEC2022-Backend/src/repository"
	"log"
)

func main() {
	if err := config.Initialize(); err != nil {
		log.Panic("init config error: ", err)
	}
	if err := logger.Initialize(); err != nil {
		log.Panic("init logger error: ", err)
	}
	logger.Info("server is initializing...")
	if err := cache.Initialize(); err != nil {
		logger.PanicAny("init cache error: ", err)
	}
	if err := repository.Initialize(); err != nil {
		logger.PanicAny("init repository error: ", err)
	}
	if err := controller.Initialize(); err != nil {
		logger.PanicAny("init controller error: ", err)
	}
	logger.Info("initialization completed, starting service...")
	if err := controller.Run(); err != nil {
		logger.PanicAny("run server error: ", err)
	}
}
