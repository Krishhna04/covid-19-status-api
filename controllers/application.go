package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"krishna/services/covid-19-status/configs"
	"krishna/services/covid-19-status/helpers"
	"krishna/services/covid-19-status/models"
)

func respond(context *gin.Context, response *models.Response) {

	context.JSON(int(response.Status), gin.H{
		configs.RESULT: response,
	})

}

func customRecover(c *gin.Context) {
	if r := recover(); r != nil {
		helpers.Logger.Error("exit crawl: PANIC occured")
		var response models.Response
		response.Data = map[string]interface{}{"error": configs.SERVER_ERROR}
		response.Status = 500
		c.JSON(int(http.StatusInternalServerError), gin.H{
			"result": response,
		})
	}
}
