package main

import (
	"krishna/services/covid-19-status/controllers"
	"krishna/services/covid-19-status/helpers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() {
	router = gin.New()

	helpers.Logger.Info("Inside Routes function")

	api_r := router.Group("/api/v1")
	api_r.POST("/search", controllers.FetchData)

}
