package main

import (
	"feriados/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.CORS())

	// Initializing controller to obtain Feriados
	feriadosController, err := controllers.NewFeriadosController()
	if err != nil {
		panic(err)
	}

	// Handling requests
	e.GET("/feriados", feriadosController.GetFeriados)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
