package controllers

import "feriados/services"

type FeriadosController struct {
	svc services.FeriadoService
}

func NewFeriadosController() FeriadosController {

}
