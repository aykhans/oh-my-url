package errors

import "errors"

var (
	ErrAPITemplateParsing = errors.New("an error occured while parsing template please try again later")
	ErrAPIInvalidURL      = errors.New("invalid URL")
	ErrAPI503             = errors.New("service unavailable please try again later")
)
