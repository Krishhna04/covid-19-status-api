package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"krishna/services/covid-19-status/models"

)

func SearchData(stringMap map[string]interface{}, input *models.Input) (map[string]interface{}, error) {

	url := "https://api.covid19api.com/country/india/status/confirmed?from=2022-10-01T00:00:00Z&to=2022-10-03T00:00:00Z"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return stringMap, err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return stringMap, err

	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return stringMap, err

	}
	fmt.Println(string(body))
	var covidResponse models.CovidResponse
	json.Unmarshal(body, &covidResponse)

	stringMap["data"]=covidResponse
	return stringMap,nil
}
