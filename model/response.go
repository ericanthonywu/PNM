package model

type (
	DefaultResponse struct {
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data"`
	}
	ErrorResponse struct {
		Message interface{} `json:"message"`
	}
	CustomResponse map[string]interface{}
	Response struct {
		Message string `json:"message"`
	}
)
