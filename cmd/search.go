package cmd

import (
	"fmt"
	"os"
	"rapide/internal/storage"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search journal entries for a keyword",
	Args:  cobra.ExactArgs(1),
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

		query := strings.ToLower(args[0])

		// Sort newest first
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Timestamp.After(entries[j].Timestamp)
		})

		foundCount := 0
		for _, e := range entries {
			contentMatch := strings.Contains(strings.ToLower(e.Content), query)
			marginMatch := strings.Contains(strings.ToLower(e.MarginKey), query)

			if contentMatch || marginMatch {
				renderEntry(e)
				foundCount++
			}
		}

		if foundCount == 0 {
			fmt.Printf("No entries found matching \"%s\"\n", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
