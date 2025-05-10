package models

import "errors"

// Custom error types for the URL shortener service
var (
	// ErrURLNotFound is returned when a URL is not found in the database
	ErrURLNotFound = errors.New("URL not found")

	// ErrInvalidURL is returned when the provided URL is invalid
	ErrInvalidURL = errors.New("invalid URL format")

	// ErrDuplicateURL is returned when trying to create a URL that already exists
	ErrDuplicateURL = errors.New("URL already exists")

	// ErrInvalidShortCode is returned when the short code is invalid
	ErrInvalidShortCode = errors.New("invalid short code")

	// ErrDatabaseError is returned when there's a database operation error
	ErrDatabaseError = errors.New("database operation failed")

	// ErrInvalidRequest is returned when the request body is invalid
	ErrInvalidRequest = errors.New("invalid request body")
)
