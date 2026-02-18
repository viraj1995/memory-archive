package models

import "time"

// Entry represents a memory archive entry
type Entry struct {
	ID        uint      `gorm:"primaryKey"`
	Content   string    `gorm:"type:text"`
	Type      string    `gorm:"index"` // personal, professional, study, etc.
	Tags      string    `gorm:"type:text"` // comma-separated tags
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName specifies the table name
func (Entry) TableName() string {
	return "entries"
}
