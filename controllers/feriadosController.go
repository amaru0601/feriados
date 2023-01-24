package controllers

import (
	"encoding/json"
	"feriados/models"
	"feriados/services"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type FeriadosController struct {
	svc services.FeriadoService
}

func NewFeriadosController() (*FeriadosController, error) {
	resp, err := http.Get("https://api.victorsanmartin.com/feriados/en.json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response models.ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &FeriadosController{
		svc: services.NewFeriadoService(response),
	}, nil
}

func (ctrl FeriadosController) GetFeriados(c echo.Context) error {
	eventType := c.QueryParam("eventType")
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")

	log.Infof("requesting GetFeriados")

	response := ctrl.svc.Feriados.Data

	if eventType != "" {
		log.Infof("filtering by type")
		response = ctrl.svc.FilterByType(eventType)
	}

	if startDate != "" && endDate != "" {
		log.Infof("filtering by date range")
		response = ctrl.svc.FilterByDateRange(startDate, endDate)
	}

	return c.JSON(http.StatusOK, response)
}
