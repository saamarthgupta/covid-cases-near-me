package model

import "time"

type CovidCasesInLocation struct {
	Location    string
	ActiveCases int
	LastUpdated time.Time
}

type CovidCasesResponse struct {
	State   CovidCasesInLocation
	Country CovidCasesInLocation
}
