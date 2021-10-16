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

// GetCovidCasesByLocation returns the number of active covid cases in user's location.
// @Summary Get the active covid cases in state and Country.
// @Description Get the active covid cases in state and Country.
// @Tags Covid Cases
// @Accept  json
// @Produce  json
// @Param latitude query string true "Latitude"
// @Param longitude query string true "Longitude"
// @Success 200 {object} model.CovidCasesResponse "Active Covid cases in State and Country"
// @Failure 404 {string} string "Unable to Fetch Data for Given Latitude and Longitudes"
// @Router /covidCases [get]
func (controller *CaseController) GetCovidCasesByLocation(c echo.Context) error {
	latitude := c.QueryParam(constants.LATITUDE)
	longitude := c.QueryParam(constants.LONGITUDE)
	userCoordinates := model.UserCoordinates{Latitude: latitude, Longitude: longitude}
	if len(latitude) == 0 || len(longitude) == 0 {
		return c.JSON(http.StatusNotFound, "Latitude and Longitudes are missing")
	}
	// Call Service to fetch Response
	response, err := controller.service.GetCasesByLocation(userCoordinates)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Unable to Fetch Data for Given Latitude and Longitudes")
	}
	return c.JSON(http.StatusOK, response)
}

// GetCaseDataFromSource fetch Covid Case Data from API and store in DB.
// @Summary fetch Covid Case Data from API and store in DB.
// @Description fetch Covid Case Data from API and store in DB.
// @Tags Covid Cases
// @Accept  json
// @Produce  json
// @Success 201 {string} string "OK"
// @Failure 404 {string} string "Resource Not Found"
// @Router /covidCases [post]
func (controller *CaseController) GetCaseDataFromSource(c echo.Context) error {
	err := controller.service.PullCaseDataFromServer()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, "OK")
}
