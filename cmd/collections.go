package cmd

import (
	"fmt"
	"os"
	"rapide/internal/storage"
	"sort"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var collectionNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8")).Bold(true).Width(15)
var countStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#757575"))

var collectionsCmd = &cobra.Command{
	Use:   "collections",
	Short: "List all unique margin keys (collections) and their item counts",
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

		counts := make(map[string]int)
		for _, e := range entries {
			margin := e.MarginKey
			if margin == "" {
				margin = "(Inbox)"
			}
			counts[margin]++
		}

		keys := make([]string, 0, len(counts))
		for k := range counts {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		if len(keys) == 0 {
			fmt.Println("No collections found.")
			return
		}

		fmt.Println("Collections:")
		fmt.Println("------------")
		for _, k := range keys {
			name := collectionNameStyle.Render(k)
			count := countStyle.Render(fmt.Sprintf("%d items", counts[k]))
			fmt.Printf("%s | %s\n", name, count)
		}
	},
}

func init() {
	rootCmd.AddCommand(collectionsCmd)
}
