package cmd

import (
	"fmt"
	"os"
	"rapide/internal/model"
	"rapide/internal/storage"
	"time"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [id]",
	Short: "Migrate a task to today (>)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error initializing storage: %v\n", err)
			os.Exit(1)
		}

		id := args[0]
		entries, err := s.List()
		if err != nil {
			fmt.Printf("Error reading entries: %v\n", err)
			os.Exit(1)
		}

		var targetEntry model.Entry
		found := false
		for _, e := range entries {
			if e.ID == id {
				targetEntry = e
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Entry with ID %s not found\n", id)
			os.Exit(1)
		}

		// 1. Mark old entry as migrated (>)
		targetEntry.Bullet = ">"
		if err := s.Update(id, targetEntry); err != nil {
			fmt.Printf("Error updating entry: %v\n", err)
			os.Exit(1)
		}

		// 2. Create new entry for today (•)
		newEntry := model.Entry{
			Timestamp: time.Now(),
			MarginKey: targetEntry.MarginKey,
			Bullet:    "•",
			Content:   targetEntry.Content,
			Priority:  targetEntry.Priority,
		}

		newID, err := s.Append(newEntry)
		if err != nil {
			fmt.Printf("Error creating new entry: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s Task %s migrated to today (New ID: %s).\n", successStyle.Render("✓"), id, newID)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
