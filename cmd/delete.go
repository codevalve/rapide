package cmd

import (
	"fmt"
	"os"
	"rapide/internal/storage"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a journal entry by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error initializing storage: %v\n", err)
			os.Exit(1)
		}

		id := args[0]
		if err := s.Delete(id); err != nil {
			fmt.Printf("Error deleting entry: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s Entry %s deleted.\n", successStyle.Render("✓"), id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
