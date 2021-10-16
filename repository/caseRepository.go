package repository

import "covid_cases_near_me/config"

type CaseRepository struct{

}

func NewCaseRepository(config *config.Config) *CaseRepository{
	return &CaseRepository{}
}
