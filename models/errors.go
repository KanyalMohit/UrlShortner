package models

import "errors"

var (
	ErrURLNotFound      = errors.New("URL not found")
	ErrInvalidURL       = errors.New("invalid URL format")
	ErrDuplicateURL     = errors.New("URL already exists")
	ErrInvalidShortCode = errors.New("invalid short code")
)
