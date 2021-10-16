package main

import (
	"covid_cases_near_me/router"
	"os"

	_ "covid_cases_near_me/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	router.INIT(e)
	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "11111" // Default port if not specified
	}
	e.Logger.Fatal(e.Start(":" + port))
}
