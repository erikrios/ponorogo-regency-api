package service

import (
	"errors"

	"github.com/erikrios/ponorogo-regency-api/repository"
)

var (
	ErrDataNotFound = errors.New("service: data with given params not found")
	ErrRepository   = errors.New("service: repository error happened")
)

func mapError(from error) error {
	if errors.Is(from, repository.ErrQueryNotFound) {
		return ErrDataNotFound
	} else {
		return ErrRepository
	}
}
