package rpc

import (
	"errors"

	"github.com/erikrios/ponorogo-regency-api/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(from error) error {
	var code codes.Code
	var message string
	if errors.Is(from, service.ErrDataNotFound) {
		code = codes.NotFound
		message = "Resource with given ID not found."
	} else if errors.Is(from, service.ErrRepository) {
		code = codes.Internal
		message = "Something went wrong."
	} else {
		code = codes.Internal
		message = "Something went wrong."
	}

	return status.Errorf(code, message)
}
