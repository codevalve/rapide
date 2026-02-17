package cmd

import (
	"fmt"
	"os"
	"rapide/internal"
	"rapide/internal/storage"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [id] [new content]",
	Short: "Edit a journal entry by ID",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error initializing storage: %v\n", err)
			os.Exit(1)
		}

		id := args[0]
		newEntry := internal.ParseEntry(args[1:])

		if err := s.Update(id, newEntry); err != nil {
			fmt.Printf("Error updating entry: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s Entry %s updated.\n", successStyle.Render("✓"), id)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}
