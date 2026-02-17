package cmd

import (
	"fmt"
	"os"
	"rapide/internal/storage"
	"sort"

	"github.com/spf13/cobra"
)

var unfinishedCmd = &cobra.Command{
	Use:   "unfinished",
	Short: "List all pending tasks (•)",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error initializing storage: %v\n", err)
			os.Exit(1)
		}

		entries, err := s.List()
		if err != nil {
			fmt.Printf("Error reading entries: %v\n", err)
			os.Exit(1)
		}

		// Sort newest first
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Timestamp.After(entries[j].Timestamp)
		})

		foundCount := 0
		for _, e := range entries {
			if e.Bullet == "•" {
				renderEntry(e)
				foundCount++
			}
		}

		if foundCount == 0 {
			fmt.Println("No pending tasks found.")
		}
	},
}

func init() {
	rootCmd.AddCommand(unfinishedCmd)
}
