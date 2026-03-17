package cmd

import (
	"fmt"
	"os"
	"rapide/internal"
	"rapide/internal/storage"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var Version = "3.0.2"

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
		// If the first argument is a registered subcommand, let it handle the execution.
		// Cobra usually does this automatically, but since we have a greedy Run function
		// on the root command, we need to be careful not to consume subcommands.
		if len(args) > 0 {
			for _, sub := range cmd.Commands() {
				if sub.Name() == args[0] {
					// Found a subcommand match, this Run shouldn't have been called
					// Or we can manually execute it if needed, but usually we just return
					// and let Cobra handle it if Run was NOT defined.
					// However, since Run IS defined, it takes precedence.
					sub.Run(sub, args[1:])
					return
				}
			}
		}

		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error initializing storage: %v\n", err)
			os.Exit(1)
		}

		// First-run hint for empty journal
		if existing, _ := s.List(); len(existing) == 0 {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render(
				"Tip: Run 'rapide init' for an interactive setup wizard and example entries."))
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
