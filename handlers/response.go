package handlers

import (
	"log"

	"github.com/gotired/golang_backend/models"
)

type Success struct {
}

func (s Success) Detail(details string) models.DetailSchema {
	return models.DetailSchema{
		Status:   "success",
		Detail:   details,
		Location: nil,
	}
}

func (s Success) Data(data any) models.DataSchema {
	return models.DataSchema{
		Status: "success",
		Data:   data,
	}
}

type Failure struct{}

func (s Failure) Detail(details string, location string) models.DetailSchema {
	var loc *string
	if location != "" {
		loc = &location
	}
	log.Fatalf("Error details : %s", details)
	return models.DetailSchema{
		Status:   "failure",
		Detail:   details,
		Location: loc,
	}
}

func (s Failure) Data(data any) models.DataSchema {
	log.Fatalf("Error data : %s", data)
	return models.DataSchema{
		Status: "failure",
		Data:   data,
	}
}
