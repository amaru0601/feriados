package services

import (
	"feriados/models"
	"testing"
	"time"
)

func TestFilterByType(t *testing.T) {
	feriadoService := FeriadoService{
		Feriados: models.ApiResponse{
			Data: []models.Data{
				{Type: "religioso"},
				{Type: "civil"},
				{Type: "religioso"},
			},
		},
	}

	// Test filtering by religioso
	filteredData := feriadoService.FilterByType("religioso")
	if len(filteredData) != 2 {
		t.Errorf("Expected 2 religioso events, got %d", len(filteredData))
	}
	for _, data := range filteredData {
		if data.Type != "religioso" {
			t.Errorf("Expected event of type religioso, got %s", data.Type)
		}
	}

	// Test filtering by civil
	filteredData = feriadoService.FilterByType("civil")
	if len(filteredData) != 1 {
		t.Errorf("Expected 1 civil event, got %d", len(filteredData))
	}
	for _, data := range filteredData {
		if data.Type != "civil" {
			t.Errorf("Expected event of type civil, got %s", data.Type)
		}
	}
}

func TestFilterByDateRange(t *testing.T) {
	feriadoService := FeriadoService{
		Feriados: models.ApiResponse{
			Data: []models.Data{
				{Date: "2022-01-01", Type: "religioso"},
				{Date: "2022-02-01", Type: "civil"},
				{Date: "2022-03-01", Type: "religioso"},
				{Date: "2022-04-01", Type: "civil"},
				{Date: "2022-05-01", Type: "religioso"},
			},
		},
	}

	// Test filtering by date range
	startDate := "2022-02-01"
	endDate := "2022-04-01"
	filteredData, _ := feriadoService.FilterByDateRange(startDate, endDate)
	if len(filteredData) != 3 {
		t.Errorf("Expected 3 events within the date range, got %d", len(filteredData))
	}
	for _, data := range filteredData {
		date, _ := time.Parse("2006-01-02", data.Date)
		start, _ := time.Parse("2006-01-02", startDate)
		end, _ := time.Parse("2006-01-02", endDate)
		if !(date.After(start) || date.Equal(start)) && !(date.Before(end) || date.Equal(end)) {
			t.Errorf("Expected event within the date range, got %s", data.Date)
		}
	}
}
