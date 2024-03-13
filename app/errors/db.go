package errors

import "errors"

var (
	ErrDBCreateURL   = errors.New("an error occured please try again later")
	ErrDBGetURL      = errors.New("an error occured please try again later")
	ErrDBURLNotFound = errors.New("url not found")
)
