package service

import (
	"covid_cases_near_me/model"
	"covid_cases_near_me/repository"
)

type CaseService struct{
	caseRepository *repository.CaseRepository
}

func NewCaseService(repository *repository.CaseRepository) *CaseService{
	return &CaseService{caseRepository: repository}
}

func (caseService *CaseService) GetCasesByLocation(coordinates model.UserCoordinates) (model.CovidCasesResponse, error) {
	return model.CovidCasesResponse{}, nil
}

func (caseService *CaseService) PullCaseDataFromServer() error{
	return nil
}
