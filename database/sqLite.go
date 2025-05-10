package database

import (
	"database/sql"
	"log"
	"urlShortener/models"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDB struct {
	db *sql.DB
}

func NewSQLiteDB(dbPath string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected to SQLite database")
	return &SQLiteDB{db: db}, nil
}

func (s *SQLiteDB) InitSchema() error {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			original_url TEXT NOT NULL,
			short_code TEXT NOT NULL UNIQUE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS idx_short_code ON urls(short_code);
		CREATE INDEX IF NOT EXISTS idx_original_url ON urls(original_url);
	`
	_, err := s.db.Exec(createTableSQL)
	return err
}

func (s *SQLiteDB) Close() error {
	return s.db.Close()
}

// CreateURL stores a new URL mapping
func (s *SQLiteDB) CreateURL(originalURL, shortCode string) (*models.URL, error) {
	query := `
		INSERT INTO urls (original_url, short_code, created_at, updated_at)
		VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id, created_at, updated_at
	`

	var url models.URL
	err := s.db.QueryRow(query, originalURL, shortCode).Scan(
		&url.ID,
		&url.CreatedAt,
		&url.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	url.OriginalURL = originalURL
	url.ShortCode = shortCode

	return &url, nil
}

// GetURLByShortCode retrieves a URL by its short code
func (s *SQLiteDB) GetURLByShortCode(shortCode string) (*models.URL, error) {
	query := `
		SELECT id, original_url, short_code, created_at, updated_at
		FROM urls
		WHERE short_code = ?
	`

	var url models.URL
	err := s.db.QueryRow(query, shortCode).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.CreatedAt,
		&url.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, models.ErrURLNotFound
	}
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// GetURLByOriginalURL retrieves a URL by its original URL
func (s *SQLiteDB) GetURLByOriginalURL(originalURL string) (*models.URL, error) {
	query := `
		SELECT id, original_url, short_code, created_at, updated_at
		FROM urls
		WHERE original_url = ?
	`

	var url models.URL
	err := s.db.QueryRow(query, originalURL).Scan(
		&url.ID,
		&url.OriginalURL,
		&url.ShortCode,
		&url.CreatedAt,
		&url.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, models.ErrURLNotFound
	}
	if err != nil {
		return nil, err
	}

	return &url, nil
}

// ListURLs returns a paginated list of URLs
func (s *SQLiteDB) ListURLs(limit, offset int) ([]*models.URL, error) {
	query := `
		SELECT id, original_url, short_code, created_at, updated_at
		FROM urls
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var urls []*models.URL
	for rows.Next() {
		var url models.URL
		err := rows.Scan(
			&url.ID,
			&url.OriginalURL,
			&url.ShortCode,
			&url.CreatedAt,
			&url.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		urls = append(urls, &url)
	}

	return urls, nil
}

// DeleteURL deletes a URL by its short code
func (s *SQLiteDB) DeleteURL(shortCode string) error {
	query := `
		DELETE FROM urls
		WHERE short_code = ?
	`

	result, err := s.db.Exec(query, shortCode)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return models.ErrURLNotFound
	}

	return nil
}
