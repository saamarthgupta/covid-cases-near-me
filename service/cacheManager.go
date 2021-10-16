package service

import (
	"covid_cases_near_me/config"
	"covid_cases_near_me/constants"
	"covid_cases_near_me/model"
	"covid_cases_near_me/repository"
	"encoding/json"
	"log"
	"strings"
)

type CacheManager struct {
	caseRepository *repository.CaseRepository
	redisObject    *config.RedisCache
}

func NewCacheManager(repository *repository.CaseRepository) *CacheManager {
	return &CacheManager{caseRepository: repository, redisObject: config.NewRedisObject()}
}

func (cacheManager *CacheManager) GetCovidCasesByLocation(state string, country string) (model.CovidCasesResponse, error) {

	stateData, err1 := cacheManager.getCasesFromCache(state, constants.STATE_DATA_CACHE_KEY)
	countryData, err2 := cacheManager.getCasesFromCache(country, constants.COUNTRY_DATA_CACHE_KEY)
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

func (cacheManager *CacheManager) getCasesFromCache(place string, cacheKeyPrefix string) (model.CaseData, error) {
	rawData, err := cacheManager.redisObject.Get(cacheKeyPrefix + strings.ReplaceAll(place, " ", ""))
	placeData := model.CaseData{}
	err1 := err
	if err != nil {
		if cacheKeyPrefix == constants.STATE_DATA_CACHE_KEY {
			placeData, err1 = cacheManager.caseRepository.GetCasesByState(place)
		} else {
			placeData, err1 = cacheManager.caseRepository.GetCasesByCountry(place)
		}
		if err1 == nil {
			jsonData, errorr := json.Marshal(placeData)
			if errorr != nil {
				log.Println(err)
			}
			err1 = cacheManager.redisObject.Set(cacheKeyPrefix+strings.ReplaceAll(place, " ", ""), string(jsonData), constants.CACHE_TTL)
		}
	} else {
		jsonErr := json.Unmarshal([]byte(rawData), &placeData)
		if jsonErr != nil {
			log.Println("Error Decoding StateData fetched from Redis")
		}
	}
	return placeData, err1
}
