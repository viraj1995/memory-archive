package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/yourusername/memarc/internal/cli"
	"github.com/yourusername/memarc/internal/storage"
)

var dbPath string
var db *storage.DB

var rootCmd = &cobra.Command{
	Use:   "memarc",
	Short: "Personal memory archive CLI",
	Long:  "A simple CLI tool to store, retrieve, and organize your memories",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		db, err = storage.New(dbPath)
		if err != nil {
			log.Fatalf("failed to initialize database: %v", err)
		}
	},
}

func init() {
	// Set default DB path
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	dbPath = filepath.Join(home, ".memarc", "memarc.db")

	// Root flags
	rootCmd.PersistentFlags().StringVar(&dbPath, "db", dbPath, "path to database file")

	// Create wrapper functions that use the db instance
	rootCmd.AddCommand(
		addCmd(),
		listCmd(),
	)
}

func addCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [content]",
		Short: "Add a new memory entry",
		Long:  "Add a new memory entry with optional type and tags",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.Add(db, cmd, args)
		},
	}
	cmd.Flags().StringP("type", "t", "personal", "entry type (personal, professional, study, etc.)")
	cmd.Flags().StringP("tags", "g", "", "comma-separated tags")
	return cmd
}

func listCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all memory entries",
		Long:  "Display all stored memory entries with timestamps",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cli.List(db)
		},
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
