package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/yourusername/memarc/internal/models"
	"github.com/yourusername/memarc/internal/storage"
)

// Add handles adding a new memory entry
func Add(db *storage.DB, cmd *cobra.Command, args []string) error {
	content := args[0]
	entryType, _ := cmd.Flags().GetString("type")
	tags, _ := cmd.Flags().GetString("tags")

	entry := &models.Entry{
		Content:   content,
		Type:      entryType,
		Tags:      tags,
		CreatedAt: time.Now(),
	}

	if err := db.CreateEntry(entry); err != nil {
		return fmt.Errorf("failed to create entry: %w", err)
	}

	fmt.Printf("✓ Entry added (ID: %d)\n", entry.ID)
	return nil
}
