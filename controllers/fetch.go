package controllers

import (
	"krishna/services/covid-19-status/configs"
	"krishna/services/covid-19-status/helpers"
	"krishna/services/covid-19-status/models"
	"krishna/services/covid-19-status/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchData(context *gin.Context) {
	defer customRecover(context)

	stringMap := make(map[string]interface{})
	response := models.Response{}
	country := context.Request.PostFormValue(configs.COUNTRY_KEY)
	fromDate := context.Request.PostFormValue(configs.FROM_DATE_KEY)
	toDate := context.Request.PostFormValue(configs.TO_DATE_KEY)
	if country == "" || fromDate == "" || toDate == "" {
		stringMap["error"] = "Invalid input"
		response.Data = stringMap
		response.Status = http.StatusBadRequest
		respond(context, &response)
		helpers.Logger.Error(response)
		return

	}

	input := models.Input{
		Country:  country,
		FromDate: fromDate,
		ToDate:   toDate,
	}

	stringMap, err := services.SearchData(stringMap, &input)
	if err != nil {
		stringMap["err"] = err
		response.Data = stringMap
		response.Status = http.StatusInternalServerError
		respond(context, &response)
		helpers.Logger.Error(response)
		return
	}

	response.Status = http.StatusOK
	response.Data = stringMap
	respond(context, &response)

	helpers.Logger.Info(response)

}
