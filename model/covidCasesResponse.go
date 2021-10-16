package model

import "time"

type CovidCasesInLocation struct{
	location string
	cases int
}

type CovidCasesResponse struct{
	state CovidCasesInLocation
	country CovidCasesInLocation
	lastUpdated time.Time
}