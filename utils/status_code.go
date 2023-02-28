package utils

import "net/http"

func GetStatusCode(err error) int {
	var statusCode int = http.StatusInternalServerError

	if _, ok := err.(*NotValidEmailError); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*UserAlreadyExistsError); ok {
		statusCode = http.StatusConflict
	}
	return statusCode
}
