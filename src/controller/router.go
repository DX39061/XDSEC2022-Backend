package controller

import (
	"XDSEC2022-Backend/src/config"
	"github.com/gin-gonic/gin"
)

type RoutingRegisterFunc func(router *gin.RouterGroup)

var routingFunctions []RoutingRegisterFunc

func RegisterApiRoute(functions ...RoutingRegisterFunc) {
	routingFunctions = append(routingFunctions, functions...)
}

func SetupRouting(router *gin.Engine) {
	router.Use(PrepareInfo())
	apiRouter := router.Group(config.ServerConfig.ApiBasePath)

	for _, routingFunction := range routingFunctions {
		routingFunction(apiRouter)
	}
}
