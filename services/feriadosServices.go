package services

import (
	"feriados/models"
	"time"
)

type FeriadoService struct {
	Feriados models.ApiResponse
}

func NewFeriadoService(feriados models.ApiResponse) FeriadoService {
	return FeriadoService{feriados}
}

func (svc FeriadoService) FilterByType(eventType string) []models.Data {
	filteredData := make([]models.Data, 0)
	for _, d := range svc.Feriados.Data {
		if d.Type == eventType {
			filteredData = append(filteredData, d)
		}
	}
	return filteredData
}

func (svc FeriadoService) FilterByDateRange(startDate, endDate string) ([]models.Data, error) {
	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return nil, err
	}

	filteredData := make([]models.Data, 0)
	for _, d := range svc.Feriados.Data {
		date, _ := time.Parse("2006-01-02", d.Date)
		if (date.After(start) || date.Equal(start)) && (date.Before(end) || date.Equal(end)) {
			filteredData = append(filteredData, d)
		}
	}
	return filteredData, nil
}
