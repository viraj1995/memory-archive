package storage

import (
	"os"
	"path/filepath"

	"github.com/yourusername/memarc/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB wraps the gorm database connection
type DB struct {
	*gorm.DB
}

// New initializes the database and runs migrations
func New(dbPath string) (*DB, error) {
	// Create directory if it doesn't exist
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate schema
	if err := database.AutoMigrate(&models.Entry{}); err != nil {
		return nil, err
	}

	return &DB{database}, nil
}

// CreateEntry adds a new entry to the database
func (d *DB) CreateEntry(entry *models.Entry) error {
	return d.Create(entry).Error
}

// ListEntries retrieves all entries
func (d *DB) ListEntries() ([]models.Entry, error) {
	var entries []models.Entry
	err := d.Find(&entries).Error
	return entries, err
}

// GetEntry retrieves a single entry by ID
func (d *DB) GetEntry(id uint) (*models.Entry, error) {
	var entry models.Entry
	err := d.First(&entry, id).Error
	return &entry, err
}

// DeleteEntry removes an entry by ID
func (d *DB) DeleteEntry(id uint) error {
	return d.Delete(&models.Entry{}, id).Error
}
