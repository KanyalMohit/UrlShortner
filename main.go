package main

import (
	"log"
	"net/http"
	"urlShortener/api"
	"urlShortener/config"
	"urlShortener/database"
	"urlShortener/utils"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize database
	db, err := database.NewSQLiteDB(cfg.GetDatabasePath())
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	// Initialize database schema
	if err := db.InitSchema(); err != nil {
		log.Fatal("Failed to initialize database schema:", err)
	}

	// Initialize URL repository
	repo := database.NewURLRepository(db)

	// Initialize URL generator
	generator := utils.NewURLGenerator(cfg.GetShortCodeLength())

	// Initialize handler
	handler := api.NewHandler(repo, generator, cfg)

	// Setup router
	router := api.SetupRouter(handler)

	// Start server
	log.Printf("Server starting on %s", cfg.GetServerAddress())
	log.Fatal(http.ListenAndServe(cfg.GetServerAddress(), router))
}
