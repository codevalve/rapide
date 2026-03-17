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
	rootCmd.AddCommand(pinCmd)
}
