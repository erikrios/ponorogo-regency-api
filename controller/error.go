package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/erikrios/ponorogo-regency-api/service"
	"github.com/labstack/echo/v4"
)

func newErrorResponse(err error) *echo.HTTPError {
	var statusCode int
	var message string

	if errors.Is(err, service.ErrDataNotFound) {
		statusCode = http.StatusNotFound
		message = "Resource with given ID not found."
	} else if errors.Is(err, service.ErrRepository) {
		statusCode = http.StatusInternalServerError
		message = "Something went wrong."
	} else {
		statusCode = http.StatusInternalServerError
		message = "Unknows Error."
		log.Println(err)
	}

	return echo.NewHTTPError(statusCode, message)
}
