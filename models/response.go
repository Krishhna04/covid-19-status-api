package models

import "time"

type CovidResponse []struct {
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Province    string    `json:"Province"`
	City        string    `json:"City"`
	CityCode    string    `json:"CityCode"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Cases       int       `json:"Cases"`
	Status      string    `json:"Status"`
	Date        time.Time `json:"Date"`
}

type Input struct {
	Country  string `json:"country" binding:"required"`
	FromDate string `json:"from" binding:"required"`
	ToDate   string `json:"to" binding:"required"`
}

type Response struct {
	Status int16                  `json:"status"`
	Data   map[string]interface{} `json:"data"`
}
