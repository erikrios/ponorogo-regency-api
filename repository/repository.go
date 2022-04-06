package repository

import "errors"

var (
	ErrQueryNotFound = errors.New("repository: query with given params not found")
	ErrDatabase      = errors.New("repository: database query is wrong")
)
