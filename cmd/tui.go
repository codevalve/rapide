package cmd

import (
	"fmt"
	"os"
	"rapide/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch the interactive Rapanui TUI",
	Long: `Enter a state-of-the-art interactive terminal interface for Rapide.
Browse entries, filter in real-time (/), create new items (n), 
and manage your log with fast hotkeys (d, m, x, T).`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, Rapanui has crumbled: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
