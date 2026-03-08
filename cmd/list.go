package cmd

import (
	"fmt"
	"os"
	"rapide/internal/model"
	"rapide/internal/storage"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	filterWork     bool // Deprecated or unused, keeping for now
	filterMargin   string
	filterBullet   string
	filterPriority bool
	timeFilter     string
	showAll        bool
)

var listCmd = &cobra.Command{
	Use:   "list [time/margin]",
	Short: "List journal entries",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := storage.NewStorage()
		if err != nil {
			fmt.Printf("Error initializing storage: %v\n", err)
			os.Exit(1)
		}

		entries, err := s.List()
		if err != nil {
			fmt.Printf("Error reading entries: %v\n", err)
			os.Exit(1)
		}

		cfg, _ := storage.LoadConfig()

		// Sort: Pinned first, then newest first
		sort.Slice(entries, func(i, j int) bool {
			if entries[i].Pinned != entries[j].Pinned {
				return entries[i].Pinned
			}
			return entries[i].Timestamp.After(entries[j].Timestamp)
		})

		// Handle positional argument
		if len(args) > 0 {
			arg := args[0]
			// If it looks like a time filter (today or Nd), set timeFilter
			if arg == "today" || strings.HasSuffix(arg, "d") {
				if timeFilter == "" {
					timeFilter = arg
				}
			} else {
				// Otherwise treat as margin key
				if filterMargin == "" {
					filterMargin = arg
				}
			}
		}

		cutoff := time.Now().AddDate(0, 0, -30) // Default 30 days
		if timeFilter != "" {
			// Basic parsing for 3d, 7d
			if strings.HasSuffix(timeFilter, "d") {
				days := 0
				fmt.Sscanf(timeFilter, "%dd", &days)
				if days > 0 {
					cutoff = time.Now().AddDate(0, 0, -days)
				}
			} else if timeFilter == "today" {
				cutoff = time.Now().Truncate(24 * time.Hour)
			}
		}

		var filtered []model.Entry
		maxMargin := 0
		for _, e := range entries {
			// Issue #16: Auto-hide completed items
			if !showAll && cfg != nil && cfg.AutoHideDays > 0 && e.Bullet == "x" {
				hideCutoff := time.Now().Add(-time.Duration(cfg.AutoHideDays) * 24 * time.Hour)
				if e.Timestamp.Before(hideCutoff) {
					continue
				}
			}

			if e.Timestamp.Before(cutoff) {
				continue
			}

			if filterPriority && !e.Priority {
				continue
			}

			// Filter by bullet type
			if filterBullet != "" && !strings.EqualFold(e.Bullet, filterBullet) {
				continue
			}

			// Filter by margin key
			if filterMargin != "" && !strings.EqualFold(e.MarginKey, filterMargin) {
				continue
			}

			// Simple filter for margin key (fallback for timeFilter)
			if timeFilter != "" && !strings.HasSuffix(timeFilter, "d") && timeFilter != "today" {
				if !strings.EqualFold(e.MarginKey, timeFilter) {
					continue
				}
			}

			filtered = append(filtered, e)
			if len(e.MarginKey) > maxMargin {
				maxMargin = len(e.MarginKey)
			}
		}

		for _, e := range filtered {
			renderEntry(e, maxMargin)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&timeFilter, "time", "t", "", "Time filter (3d, 7d, today)")
	listCmd.Flags().StringVarP(&filterMargin, "filter", "f", "", "Filter by margin key (e.g. work)")
	listCmd.Flags().StringVarP(&filterBullet, "bullet", "b", "", "Filter by bullet type (e.g. -, O, •, x)")
	listCmd.Flags().StringVarP(&filterBullet, "type", "", "", "Alias for --bullet")
	listCmd.Flags().BoolVarP(&filterPriority, "priority", "p", false, "Filter by priority")
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all entries, including auto-hidden completed tasks")
	rootCmd.AddCommand(listCmd)
}
