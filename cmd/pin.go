package cmd

import (
	"fmt"
	"os"
	"rapide/internal/storage"

	"github.com/spf13/cobra"
)

var pinCmd = &cobra.Command{
	Use:   "pin <id>",
	Short: "Toggle pin status of an entry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		id := args[0]
		newState, err := s.TogglePin(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		status := "unpinned"
		if newState {
			status = "pinned"
		}
		fmt.Printf("Entry %s %s.\n", id, status)
	},
}

func init() {
	rootCmd.AddCommand(pinCmd)
}
