package main

import (
	"fmt"
	"krishna/services/covid-19-status/configs"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	fmt.Println("Start the service")
	SetUpRoutes()
	appPort := configs.GetEnvWithKey("APP_PORT", "")
	router.Run(appPort)
}
