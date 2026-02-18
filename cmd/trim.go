package cmd

import (
	"bufio"
	"fmt"
	"os"
	"rapide/internal/storage"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var trimBefore string

var trimCmd = &cobra.Command{
	Use:   "trim",
	Short: "Cleanup old entries by archiving or deleting them",
	Long: `Remove entries from your primary log. 
By default, it targets ALL entries. Use --before to limit the scope.
You will be prompted to either Archive them to a dated file or Delete them permanently.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		cutoff := time.Now()
		if trimBefore != "" {
			cutoff, err = time.Parse("2006-01-02", trimBefore)
			if err != nil {
				fmt.Printf("Error parsing date: %s (expected YYYY-MM-DD)\n", trimBefore)
				os.Exit(1)
			}
			// Set to start of day for the cutoff
			cutoff = time.Date(cutoff.Year(), cutoff.Month(), cutoff.Day(), 0, 0, 0, 0, cutoff.Location())
		}

		entries, err := s.List()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		var targetCount int
		for _, e := range entries {
			if e.Timestamp.Before(cutoff) {
				targetCount++
			}
		}

		if targetCount == 0 {
			fmt.Println("No entries found matching the criteria.")
			return
		}

		fmt.Printf("⚠️ This will remove %d entries from your main log.\n", targetCount)
		fmt.Print("Would you like to (a)rchive them first, (d)elete them permanently, or (c)ancel? [a/d/c]: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.ToLower(strings.TrimSpace(input))

		switch input {
		case "a", "archive", "y", "yes":
			count, filename, err := s.ArchiveBefore(cutoff)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("✓ Success! %d entries moved to %s\n", count, filename)
		case "d", "delete":
			count, err := s.TrimBefore(cutoff)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("✓ Success! %d entries permanently deleted.\n", count)
		default:
			fmt.Println("Operation cancelled.")
		}
	},
}

func init() {
	rootCmd.AddCommand(trimCmd)
	trimCmd.Flags().StringVarP(&trimBefore, "before", "b", "", "Target entries before this date (YYYY-MM-DD)")
}
