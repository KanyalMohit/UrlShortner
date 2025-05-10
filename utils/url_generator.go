package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// URLGenerator handles generation of short codes for URLs
type URLGenerator struct {
	length int
}

// NewURLGenerator creates a new URL generator with specified code length
func NewURLGenerator(length int) *URLGenerator {
	return &URLGenerator{
		length: length,
	}
}

// GenerateShortCode creates a new unique short code
func (g *URLGenerator) GenerateShortCode() (string, error) {
	// Calculate the number of bytes needed
	// We use base64 encoding which uses 6 bits per character
	// So we need (length * 6) / 8 bytes
	bytesNeeded := (g.length * 6) / 8
	if (g.length*6)%8 != 0 {
		bytesNeeded++
	}

	// Generate random bytes
	bytes := make([]byte, bytesNeeded)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Convert to base64 and trim to desired length
	code := base64.URLEncoding.EncodeToString(bytes)
	code = strings.TrimRight(code, "=") // Remove padding
	return code[:g.length], nil
}

// IsValidShortCode checks if a short code is valid
func (g *URLGenerator) IsValidShortCode(code string) bool {
	if len(code) != g.length {
		return false
	}

	// Check if code contains only valid characters
	for _, c := range code {
		if !isValidChar(c) {
			return false
		}
	}
	return true
}

// isValidChar checks if a character is valid for a short code
func isValidChar(c rune) bool {
	// Allow alphanumeric characters and some special characters
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') ||
		c == '-' || c == '_'
}
