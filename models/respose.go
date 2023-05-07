package models

type GenericResponse struct {
	Status  int64       `json:"status"`
	Code    string      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
