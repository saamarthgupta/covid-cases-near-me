package controller

import (
	"covid_cases_near_me/constants"
	"covid_cases_near_me/model"
	"covid_cases_near_me/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CaseController struct {
	service *service.CaseService
}

func NewCaseController(caseService *service.CaseService) *CaseController {
	return &CaseController{service: caseService}
}

const (
	GET_CASES_BY_LOCATION = "/covidCases"
	UPDATE_COVID_CASES    = "/covidCases"
)

func (controller *CaseController) GetCovidCasesByLocation(c echo.Context) error {
	latitude := c.QueryParam(constants.LATITUDE)
	longitude := c.QueryParam(constants.LONGITUDE)
	userCoordinates := model.UserCoordinates{Latitude: latitude, Longitude: longitude}
	if len(latitude) == 0 || len(longitude) == 0 {
		return c.JSON(http.StatusNotFound, nil)
	}
	// Call Service to fetch Response
	response, err := controller.service.GetCasesByLocation(userCoordinates)
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, response)
}

func (controller *CaseController) GetCaseDataFromSource(c echo.Context) error {
	err := controller.service.PullCaseDataFromServer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, "OK")
}
