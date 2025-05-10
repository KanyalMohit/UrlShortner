package database

import "urlShortener/models"

// URLRepositoryInterface defines the methods for URL operations
type URLRepositoryInterface interface {
	Create(originalURL string, shortCode string) (*models.URL, error)
	GetByShortCode(shortCode string) (*models.URL, error)
	GetByOriginalURL(originalURL string) (*models.URL, error)
	List(limit, offset int) ([]*models.URL, error)
	Delete(shortCode string) error
}

// URLRepository handles all database operations for URLs
type URLRepository struct {
	db *SQLiteDB
}

// NewURLRepository creates a new URL repository
func NewURLRepository(db *SQLiteDB) *URLRepository {
	return &URLRepository{db: db}
}

// Create stores a new URL mapping
func (r *URLRepository) Create(originalURL, shortCode string) (*models.URL, error) {
	// First check if URL already exists
	existingURL, err := r.GetByOriginalURL(originalURL)
	if err == nil {
		return existingURL, models.ErrDuplicateURL
	}
	if err != models.ErrURLNotFound {
		return nil, err
	}

	// Create new URL
	return r.db.CreateURL(originalURL, shortCode)
}

// GetByShortCode retrieves a URL by its short code
func (r *URLRepository) GetByShortCode(shortCode string) (*models.URL, error) {
	return r.db.GetURLByShortCode(shortCode)
}

// GetByOriginalURL retrieves a URL by its original URL
func (r *URLRepository) GetByOriginalURL(originalURL string) (*models.URL, error) {
	return r.db.GetURLByOriginalURL(originalURL)
}

// List returns a paginated list of URLs
func (r *URLRepository) List(limit, offset int) ([]*models.URL, error) {
	return r.db.ListURLs(limit, offset)
}

// Delete removes a URL by its short code
func (r *URLRepository) Delete(shortCode string) error {
	return r.db.DeleteURL(shortCode)
}
