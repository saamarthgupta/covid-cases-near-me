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
	cacheManager   *CacheManager
}

func NewCaseService(repository *repository.CaseRepository) *CaseService {
	return &CaseService{caseRepository: repository, cacheManager: NewCacheManager(repository)}
}

func (caseService *CaseService) GetCasesByLocation(coordinates model.UserCoordinates) (model.CovidCasesResponse, error) {
	country, state := GetStateAndCountry(coordinates)
	covidCasesNearUser, err := caseService.cacheManager.GetCovidCasesByLocation(state, country)
	if err != nil {
		log.Println("Error Processing Request")
		return model.CovidCasesResponse{}, err
	}
	return covidCasesNearUser, nil
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
	countryCaseData := model.CaseData{}
	if len(caseDataForCountry) > 0 {
		countryCaseData = caseDataForCountry[len(caseDataForCountry)-1]
	}

	err = caseService.caseRepository.Save(countryCaseData, stateVsCaseDataMap)

	if err != nil {
		return err
	}

	return nil
}

func getCaseDataFromApi(prefixUrl string) ([]model.CaseData, error) {
	currentTime := time.Now().UTC().Add(time.Duration(-1) * time.Hour)
	currentTimeForStart := time.Now().UTC().Add(time.Duration(-24) * time.Hour)
	startTime := time.Date(currentTimeForStart.Year(), currentTimeForStart.Month(), currentTimeForStart.Day(), 0, 0, 0, 0, currentTimeForStart.Location())
	apiSuffix := "from=" + startTime.Format(time.RFC3339) + "&to=" + currentTime.Format(time.RFC3339)
	apiUrlForStateData := prefixUrl + apiSuffix
	response, error := http.Get(apiUrlForStateData)

	if error != nil {
		return nil, error
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Println(readErr)
	}

	caseData := []model.CaseData{}
	jsonErr := json.Unmarshal(body, &caseData)

	if readErr != nil {
		log.Println(jsonErr)
	}
	return caseData, error
}
