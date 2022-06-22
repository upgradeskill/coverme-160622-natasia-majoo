package dto

type ResponseError struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
