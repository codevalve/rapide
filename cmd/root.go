package cmd

import (
	"fmt"
	"os"
	"rapide/internal"
	"rapide/internal/storage"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var Version = "dev"

var successStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#04B575")).
	Padding(0, 1)

var rootCmd = &cobra.Command{
	Use:     "rapide [margin-key] | [bullet] content",
	Short:   "Rapide is a fast CLI for Bullet Journal-style rapid logging.",
	Version: Version,
	Long: `A Go port of Rapide, designed for fast journaling.
Syntax: rapide [margin-key] | [bullet] content
Example: rapide work | - Martin updated git repo`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Default action is to add an entry
		if args[0] == "list" {
			// This will be handled by the list subcommand
			return
		}

		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error initializing storage: %v\n", err)
			os.Exit(1)
		}

		entry := internal.ParseEntry(args)
		id, err := s.Append(entry)
		if err != nil {
			fmt.Printf("Error saving entry: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s Entry added (ID: %s)\n", successStyle.Render("✓"), id)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
