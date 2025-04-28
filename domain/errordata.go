package domain

type ErrorData struct {
	Status   string      `json:"status" example:"error"`
	Code     int         `json:"code" example:"204"`
	HTTPCode int         `json:"-"`
	Message  string      `json:"message"`
	Errors   interface{} `json:"errors,omitempty"`
}
