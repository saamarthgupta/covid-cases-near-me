package service

import (
	"covid_cases_near_me/model"
	"covid_cases_near_me/repository"
)

type CacheManager struct {
	caseRepository *repository.CaseRepository
}

func NewCacheManager(repository *repository.CaseRepository) *CacheManager {
	return &CacheManager{caseRepository: repository}
}

func (cacheManager *CacheManager) GetCovidCasesByLocation(state string, country string) (model.CovidCasesResponse, error) {
	stateData, err1 := cacheManager.caseRepository.GetCasesByState(state)
	countryData, err2 := cacheManager.caseRepository.GetCasesByCountry(country)
	covidCasesInArea := model.CovidCasesResponse{}
	if err1 == nil {
		covidCasesInArea.State = model.CovidCasesInLocation{Location: stateData.Province, ActiveCases: stateData.Active, LastUpdated: stateData.Date}
	}
	if err2 == nil {
		covidCasesInArea.Country = model.CovidCasesInLocation{Location: countryData.Country, ActiveCases: countryData.Active, LastUpdated: countryData.Date}
	}
	if err1 == nil || err2 == nil {
		return covidCasesInArea, nil
	}
	if err1 == nil {
		err1 = err2
	}
	return covidCasesInArea, err1
}
