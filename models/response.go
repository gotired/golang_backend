package models

type DetailSchema struct {
	Status   string  `json:"status"`
	Detail   string  `json:"detail"`
	Location *string `json:"location,omitempty"`
}

type DataSchema struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}
