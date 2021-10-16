package model

type CaseData struct {
	Id          string `json:"ID,omitempty"`
	Country     string `json:"Country,omitempty"`
	CountryCode string `json:"CountryCode,omitempty"`
	Province    string `json:"Province,omitempty"`
	City        string `json:"City,omitempty"`
	CityCode    string `json:"CityCode,omitempty"`
	Lat         string `json:"Lat,omitempty"`
	Lon         string `json:"Lon,omitempty"`
	Confirmed   int    `json:"Confirmed,omitempty"`
	Deaths      int    `json:"Deaths,omitempty"`
	Recovered   int    `json:"Recovered,omitempty"`
	Active      int    `json:"Active,omitempty"`
	Date        string `json:"Date,omitempty"`
}
