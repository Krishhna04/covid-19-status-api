package models

import "time"

type covidResponse []struct {
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
