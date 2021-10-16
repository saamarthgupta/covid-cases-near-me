package router

import (
	"covid_cases_near_me/config"
	"covid_cases_near_me/controller"
	_ "covid_cases_near_me/docs"
	"covid_cases_near_me/repository"
	"covid_cases_near_me/service"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a covid case tracking server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host covidCases.swagger.io
// @BasePath /v1
func INIT(e *echo.Echo) {
	conf, _ := config.Load()
	caseRepository := repository.NewCaseRepository(conf)
	caseService := service.NewCaseService(caseRepository)
	caseController := controller.NewCaseController(caseService)
	e.GET(controller.GET_CASES_BY_LOCATION, func(c echo.Context) error { return caseController.GetCovidCasesByLocation(c) })
	e.POST(controller.UPDATE_COVID_CASES, func(c echo.Context) error { return caseController.GetCaseDataFromSource(c) })
	e.GET("/swagger/*", echoSwagger.WrapHandler)

}
