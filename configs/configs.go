package configs

import (
	"syscall"

	"krishna/services/covid-19-status/helpers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	if GetEnvWithKey(APP_ENVIRONMENT, DEVELOPMENT) == PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

}

func loadEnv() {
	err := godotenv.Overload()
	if err != nil {
		helpers.Logger.Fatalf("Error in loading .env file: %s", err.Error())
	}

}

func GetEnvWithKey(key string, defaultValue string) string {

	value, found := syscall.Getenv(key)
	if !found {
		syscall.Setenv(key, defaultValue)
		return defaultValue
	}
	return value
}

const (
	APP_ENVIRONMENT       string = "APP_ENVIRONMENT"
	SERVER_ERROR          string = "something went wrong in the serverside"
	TXN_ID_KEY            string = "txnid"
	DEVELOPMENT           string = "DEVELOPMENT"
	PRODUCTION            string = "PRODUCTION"
	RESULT                string = "Result"
	ERROR                 string = "Error"
	SERVICE_START_MESSAGE string = "Covid status service start..."
	MAX_API_RETRIES       int    = 2
	COUNTRY_KEY           string = "country"
	FROM_DATE_KEY         string = "from"
	TO_DATE_KEY           string = "to"
)
