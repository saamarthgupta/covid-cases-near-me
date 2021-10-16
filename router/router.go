package router

import (
	"covid_cases_near_me/config"
	"covid_cases_near_me/controller"
	"covid_cases_near_me/repository"
	"covid_cases_near_me/service"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func INIT(e *echo.Echo)  {
	conf, env := config.Load()
	caseRepository := repository.NewCaseRepository(conf)
	caseService := service.NewCaseService(caseRepository)
	caseController := controller.NewCaseController(caseService)
	e.GET(controller.GET_CASES_BY_LOCATION, func(c echo.Context) error { return caseController.GetCovidCasesByLocation(c) })
	e.POST(controller.UPDATE_COVID_CASES, func(c echo.Context) error { return caseController.GetCaseDataFromSource(c) })

	if env == config.DEV {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
}