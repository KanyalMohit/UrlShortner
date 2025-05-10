package models

import (
	"time"
)

// URL represents a shortened URL entry
type URL struct {
	ID          int64     `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortCode   string    `json:"short_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// URLRequest represents the request body for creating a new short URL
type URLRequest struct {
	OriginalURL string `json:"original_url" validate:"required,url"`
}

// URLResponse represents the response for URL operations
type URLResponse struct {
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}
