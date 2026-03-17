package cmd

import (
	"fmt"
	"os"
	"rapide/internal/storage"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as completed (x)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error initializing storage: %v\n", err)
			os.Exit(1)
		}

		id := args[0]
		entries, err := s.List()
		if err != nil {
			fmt.Printf("Error reading entries: %v\n", err)
			os.Exit(1)
		}

		found := false
		for _, e := range entries {
			if e.ID == id {
				e.Bullet = "x"
				if err := s.Update(id, e); err != nil {
					fmt.Printf("Error updating entry: %v\n", err)
					os.Exit(1)
				}
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Entry with ID %s not found\n", id)
			os.Exit(1)
		}

		fmt.Printf("%s Task %s marked as done.\n", successStyle.Render("✓"), id)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		s, err := storage.NewStorage()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		ids, err := s.GetRecentIDs()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}

		return ids, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
