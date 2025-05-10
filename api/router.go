package api

import (
	"github.com/gorilla/mux"
)

// SetupRouter configures all the routes for our API
func SetupRouter(handler *Handler) *mux.Router {
	router := mux.NewRouter()

	// API routes
	router.HandleFunc("/api/urls", handler.CreateShortURL).Methods("POST")
	router.HandleFunc("/api/urls/{shortCode}", handler.GetURLInfo).Methods("GET")
	router.HandleFunc("/api/urls/{shortCode}", handler.DeleteURL).Methods("DELETE")

	// Redirect route (this should be the last route)
	router.HandleFunc("/{shortCode}", handler.RedirectToOriginalURL).Methods("GET")

	return router
}
