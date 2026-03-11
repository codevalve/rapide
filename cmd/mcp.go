package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"rapide/internal/mcp"
	"rapide/internal/storage"
	"syscall"

	"github.com/spf13/cobra"
)

var mcpCmd = &cobra.Command{
	Use:    "mcp",
	Short:  "MCP server commands",
	Hidden: true,
}

var mcpStartCmd = &cobra.Command{
	Use:    "start",
	Short:  "Start the MCP server",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := storage.NewStorage()
		if err != nil {
			return fmt.Errorf("failed to initialize storage: %w", err)
		}

		adapter := mcp.NewJournalAdapter(s)
		server := mcp.NewServer(adapter)

		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		fmt.Fprintln(os.Stderr, "Starting Rapide MCP server (stdio)...")
		return server.Start(ctx)
	},
}

func init() {
	rootCmd.AddCommand(mcpCmd)
	mcpCmd.AddCommand(mcpStartCmd)
}
