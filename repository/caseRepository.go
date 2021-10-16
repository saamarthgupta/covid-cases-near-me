package repository

import (
	"covid_cases_near_me/config"
	"covid_cases_near_me/model"
)

type CaseRepository struct {
}

func NewCaseRepository(config *config.Config) *CaseRepository {
	return &CaseRepository{}
}

func (caseRepository *CaseRepository) Save(caseData model.CaseData, stateVsCaseDataMap map[string]model.CaseData) {

}
