package controller

import (
	"github.com/gin-gonic/gin"
	"xdsec-join/src/config"
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
