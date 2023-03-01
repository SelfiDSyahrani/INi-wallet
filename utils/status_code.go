package utils

import "net/http"

func GetStatusCode(err error) int {
	var statusCode int = http.StatusInternalServerError

	if _, ok := err.(*NotValidEmailError); ok {
		statusCode = http.StatusUnprocessableEntity
	} else if _, ok := err.(*UserAlreadyExistsError); ok {
		statusCode = http.StatusConflict
	} else if _, ok := err.(*IncorrectCredentialsError); ok {
		statusCode = http.StatusUnauthorized
	} else if _, ok := err.(*UserNotFoundError); ok {
		statusCode = http.StatusNotFound
	}
	return statusCode
}
