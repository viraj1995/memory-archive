package cli

import (
	"fmt"
	"strings"

	"github.com/yourusername/memarc/internal/storage"
)

// List displays all memory entries
func List(db *storage.DB) error {
	entries, err := db.ListEntries()
	if err != nil {
		return fmt.Errorf("failed to retrieve entries: %w", err)
	}

	if len(entries) == 0 {
		fmt.Println("No entries found")
		return nil
	}

	fmt.Println("\n=== Memory Archive ===")
	for _, entry := range entries {
		fmt.Printf("[%d] %s (%s)\n", entry.ID, entry.CreatedAt.Format("2006-01-02 15:04"), entry.Type)
		fmt.Printf("    %s\n", truncate(entry.Content, 60))
		if entry.Tags != "" {
			fmt.Printf("    Tags: %s\n", entry.Tags)
		}
		fmt.Println()
	}
	return nil
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return strings.TrimSpace(s[:maxLen]) + "..."
}
