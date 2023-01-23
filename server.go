package main

import (
	"feriados/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.CORS())

	feriadosController, err := controllers.NewFeriadosController()
	if err != nil {
		panic(err)
	}

	e.GET("/feriados", feriadosController.GetFeriados)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
