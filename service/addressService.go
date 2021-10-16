package service

import (
	"covid_cases_near_me/constants"
	"covid_cases_near_me/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ReverseGeoCodingResponse struct {
	PlaceId     string          `json:"place_id,omitempty"`
	License     string          `json:"licence,omitempty"`
	OsmType     string          `json:"osm_type,omitempty"`
	OsmId       string          `json:"osm_id,omitempty"`
	Lat         string          `json:"lat,omitempty"`
	Lon         string          `json:"lon,omitempty"`
	DisplayName string          `json:"display_name,omitempty"`
	Address     AddressResponse `json:"address,omitempty"`
	BoundingBox []string        `json:"boundingbox,omitempty"`
}

type AddressResponse struct {
	County        string `json:"county,omitempty"`
	StateDistrict string `json:"state_district,omitempty"`
	State         string `json:"state,omitempty"`
	Postcode      string `json:"postcode,omitempty"`
	Country       string `json:"country,omitempty"`
	CountryCode   string `json:"country_code,omitempty"`
	Road          string `json:"road,omitempty"`
	Town          string `json:"town,omitempty"`
	Village       string `json:"village,omitempty"`
}

func GetStateAndCountry(coordinates model.UserCoordinates) (state string, country string) {
	country, state = getCountryAndStateInfo(coordinates)
	return country, state
}

func getCountryAndStateInfo(coordinates model.UserCoordinates) (state string, country string) {
	reverseGeoCodingResponse := reverseGeoCode(coordinates)
	state, country = parseGeoCodingResponse(reverseGeoCodingResponse)
	return country, state
}

func reverseGeoCode(coordinates model.UserCoordinates) ReverseGeoCodingResponse {
	reverseGeoCodingResponse := ReverseGeoCodingResponse{}
	apiKey := "pk.d2a4dc12bab8b79dc30109f8d3e58d11"
	apiUrl := constants.REVERSE_GEOCODING_BASE_URL + apiKey + "&lat=" + coordinates.Latitude + "&lon=" + coordinates.Longitude + "&format=json"
	response, error := http.Get(apiUrl)
	if error != nil {
		return reverseGeoCodingResponse
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Println(readErr)
	}

	jsonErr := json.Unmarshal(body, &reverseGeoCodingResponse)
	if readErr != nil {
		log.Println(jsonErr)
	}
	return reverseGeoCodingResponse
}

func parseGeoCodingResponse(reverseGeoCodingResponse ReverseGeoCodingResponse) (state string, country string) {
	return reverseGeoCodingResponse.Address.State, reverseGeoCodingResponse.Address.Country
}
