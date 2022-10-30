package configs

import (
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"krishna/services/covid-19-status/helpers"
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
	APP_ENVIRONMENT              string        = "APP_ENVIRONMENT"
	AWS_REGION_NAME              string        = "ap-south-1"
	SERVICE_ID                   int16         = 375
	SERVER_ERROR                 string        = "something went wrong in the serverside"
	FAST_STORE_DATASTRUCTURE_SET string        = "HSET"
	FAST_STORE_DATASTRUCTURE_GET string        = "HGET"
	TXN_ID_KEY                   string        = "txnid"
	DEVELOPMENT                  string        = "DEVELOPMENT"
	PRODUCTION                   string        = "PRODUCTION"
	RESULT                       string        = "Result"
	ERROR                        string        = "Error"
	SERVICE_START_MESSAGE        string        = "CA-VERIFICATION service start..."
	MAX_API_RETRIES              int           = 2
	GST_IN                       string        = "GST_IN"
	FINANCIAL_YEAR               string        = "FINANCIAL_YEAR"
	SAMPLE_GSTIN                 string        = "06AAHCA4013B1ZG"
	SAMPLE_FINANCIAL_YEAR        string        = "2019-20"
	TESTING_PDF_FILEPATH         string        = "./output.pdf"
	PRESIGN_TIME                 time.Duration = 15 * time.Minute
	SEARCH_TAXPAYER_API_URL      string        = "SEARCH_TAXPAYER_API_URL"
	VIEW_RETURNS_API_URL         string        = "VIEW_RETURNS_API_URL"
	OUTPUT_KEY                   string        = "output.pdf"
)
