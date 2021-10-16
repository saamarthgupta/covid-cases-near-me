package service

import (
	"covid_cases_near_me/constants"
	"covid_cases_near_me/model"
	"covid_cases_near_me/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CaseService struct {
	caseRepository *repository.CaseRepository
}

func NewCaseService(repository *repository.CaseRepository) *CaseService {
	return &CaseService{caseRepository: repository}
}

func (caseService *CaseService) GetCasesByLocation(coordinates model.UserCoordinates) (model.CovidCasesResponse, error) {
	state, country := GetStateAndCountry(coordinates)
	print(state, country)
	return model.CovidCasesResponse{}, nil
}

func (caseService *CaseService) PullCaseDataFromServer() error {
	caseDataForState, err := getCaseDataFromApi(constants.FETCH_CURRENT_CASES_URL)
	if err != nil {
		return err
	}
	stateVsCaseDataMap := make(map[string]model.CaseData)
	for _, element := range caseDataForState {
		stateVsCaseDataMap[element.Province] = element
	}

	caseDataForCountry, errCountry := getCaseDataFromApi(constants.FETCH_CURRENT_CASES_BY_COUNTRY)
	if errCountry != nil {
		return errCountry
	}
	countryCaseData := caseDataForCountry[len(caseDataForCountry)-1]
	caseService.caseRepository.Save(countryCaseData, stateVsCaseDataMap)
	return nil
}

func getCaseDataFromApi(prefixUrl string) ([]model.CaseData, error) {
	currentTime := time.Now().UTC().Add(time.Duration(-1) * time.Hour)
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	apiSuffix := "from=" + startTime.Format(time.RFC3339) + "&to=" + currentTime.Format(time.RFC3339)
	apiUrlForStateData := prefixUrl + apiSuffix
	response, error := http.Get(apiUrlForStateData)

	if error != nil {
		return nil, error
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	caseData := []model.CaseData{}
	jsonErr := json.Unmarshal(body, &caseData)

	if readErr != nil {
		log.Fatal(jsonErr)
	}
	return caseData, error
}
