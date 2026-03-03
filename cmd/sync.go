package cmd

import (
	"fmt"
	"os"
	"rapide/internal/storage"

	"github.com/spf13/cobra"
)

var setupURL string
var autoSyncFlag bool

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync items with a private git repo",
	Long:  `Sync your rapid log with a configured remote Git repository using pull (rebase) and push.`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error initializing storage: %v\n", err)
			os.Exit(1)
		}

		if setupURL != "" {
			fmt.Printf("Setting up git remote: %s...\n", setupURL)
			if err := s.SetupGit(setupURL); err != nil {
				fmt.Printf("Error during setup: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Setup successful!")
		}

		// Handle --autosync flag independently of --setup
		if cmd.Flags().Changed("autosync") {
			cfg, _ := storage.LoadConfig()
			cfg.AutoSync = autoSyncFlag
			if err := storage.SaveConfig(cfg); err != nil {
				fmt.Printf("Error saving config: %v\n", err)
				os.Exit(1)
			}
			state := "disabled"
			if autoSyncFlag {
				state = "enabled"
			}
			fmt.Printf("Autosync %s.\n", state)
		}

		// If --setup was the only thing, don't run a full sync immediately unless requested
		// or if called without flags
		if setupURL == "" && !cmd.Flags().Changed("autosync") || len(args) == 0 && setupURL == "" {
			fmt.Println("Syncing...")
			if err := s.Sync(); err != nil {
				fmt.Printf("Error during sync: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Sync complete!")
		}
	},
}

func init() {
	syncCmd.Flags().StringVarP(&setupURL, "setup", "s", "", "Configure the remote Git repository URL")
	syncCmd.Flags().BoolVarP(&autoSyncFlag, "autosync", "a", false, "Enable or disable automatic sync (true/false)")
	rootCmd.AddCommand(syncCmd)
}
