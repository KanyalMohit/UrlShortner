package api

import (
	"encoding/json"
	"net/http"
	"urlShortener/config"
	"urlShortener/database"
	"urlShortener/models"
	"urlShortener/utils"

	"github.com/gorilla/mux"
)

// Handler handles all HTTP requests
type Handler struct {
	repo      *database.URLRepository
	generator *utils.URLGenerator
	config    *config.Config
}

// NewHandler creates a new handler instance
func NewHandler(repo *database.URLRepository, generator *utils.URLGenerator, config *config.Config) *Handler {
	return &Handler{
		repo:      repo,
		generator: generator,
		config:    config,
	}
}

// CreateShortURL handles the creation of a new short URL
func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req models.URLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Generate short code
	shortCode, err := h.generator.GenerateShortCode()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to generate short code")
		return
	}

	// Create URL in database
	url, err := h.repo.Create(req.OriginalURL, shortCode)
	if err != nil {
		if err == models.ErrDuplicateURL {
			respondWithError(w, http.StatusConflict, "URL already exists")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to create short URL")
		return
	}

	// Create response
	response := models.URLResponse{
		ShortURL:    h.config.GetBaseURL() + "/" + url.ShortCode,
		OriginalURL: url.OriginalURL,
		CreatedAt:   url.CreatedAt,
	}

	respondWithJSON(w, http.StatusCreated, response)
}

// RedirectToOriginalURL handles redirection to the original URL
func (h *Handler) RedirectToOriginalURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// Validate short code
	if !h.generator.IsValidShortCode(shortCode) {
		respondWithError(w, http.StatusBadRequest, "Invalid short code")
		return
	}

	// Get URL from database
	url, err := h.repo.GetByShortCode(shortCode)
	if err != nil {
		if err == models.ErrURLNotFound {
			respondWithError(w, http.StatusNotFound, "URL not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve URL")
		return
	}

	// Redirect to original URL
	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}

// GetURLInfo returns information about a short URL
func (h *Handler) GetURLInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// Validate short code
	if !h.generator.IsValidShortCode(shortCode) {
		respondWithError(w, http.StatusBadRequest, "Invalid short code")
		return
	}

	// Get URL from database
	url, err := h.repo.GetByShortCode(shortCode)
	if err != nil {
		if err == models.ErrURLNotFound {
			respondWithError(w, http.StatusNotFound, "URL not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve URL")
		return
	}

	// Create response
	response := models.URLResponse{
		ShortURL:    h.config.GetBaseURL() + "/" + url.ShortCode,
		OriginalURL: url.OriginalURL,
		CreatedAt:   url.CreatedAt,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// DeleteURL deletes a short URL
func (h *Handler) DeleteURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	// Validate short code
	if !h.generator.IsValidShortCode(shortCode) {
		respondWithError(w, http.StatusBadRequest, "Invalid short code")
		return
	}

	// Delete URL from database
	err := h.repo.Delete(shortCode)
	if err != nil {
		if err == models.ErrURLNotFound {
			respondWithError(w, http.StatusNotFound, "URL not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "Failed to delete URL")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Helper function to respond with JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Helper function to respond with error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.ErrorResponse{Error: message})
}
