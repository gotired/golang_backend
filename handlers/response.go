package handlers

import (
	"github.com/gotired/golang_backend/models"
)

type Success struct{}

func (s Success) Detail(details string) models.DetailSchema {
	response := models.DetailSchema{
		Status:   "success",
		Detail:   details,
		Location: nil,
	}
	return response
}

func (s Success) Data(data any) models.DataSchema {
	response := models.DataSchema{
		Status: "success",
		Data:   data,
	}
	return response
}

type Failure struct{}

func (s Failure) Detail(details, location string) models.DetailSchema {
	var locPtr *string
	if location != "" {
		locPtr = &location
	}
	response := models.DetailSchema{
		Status:   "failure",
		Detail:   details,
		Location: locPtr,
	}
	return response
}

func (s Failure) Data(data any) models.DataSchema {
	response := models.DataSchema{
		Status: "failure",
		Data:   data,
	}
	return response
}
