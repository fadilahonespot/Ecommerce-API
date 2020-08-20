package model

type Respons struct {
	Success bool        `json:"Success" example:"true"`
	Message string      `json:"Message" example:"Success"`
	Data    interface{} `json:"data,omitempty"`
}

type ResponsFalse struct {
	Success bool        `json:"Success" example:"false"`
	Message string      `json:"Message" example:"message"`
	Data    interface{} `json:"data,omitempty"`
}
