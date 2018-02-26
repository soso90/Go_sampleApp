package main

import (
	"Go_sampleApp/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS()) // CORS
	// static
	e.Static("/static", "static")
	// controller initialaize
	// MainController
	controllers.MainController{}.Init(e.Group("/"))

	// Server
	e.Start(":1323")

	//e.Logger.Fatal(e.Start(":1323"))
}
