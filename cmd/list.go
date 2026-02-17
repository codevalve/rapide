package cmd

import (
	"fmt"
	"os"
	"rapide/internal/storage"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	filterWork     bool // Deprecated or unused, keeping for now
	filterMargin   string
	filterPriority bool
	timeFilter     string
)

var (
	timestampStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#757575"))
	marginStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8")).Bold(true).Width(12)
	bulletStyle    = lipgloss.NewStyle().Bold(true)
	priorityStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
	noteStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#EEEEEE"))
	idStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#555555")).Italic(true).Width(5)
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List journal entries",
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

		// Sort newest first
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Timestamp.After(entries[j].Timestamp)
		})

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

		for _, e := range entries {
			if e.Timestamp.Before(cutoff) {
				continue
			}

			if filterPriority && !e.Priority {
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

			ts := timestampStyle.Render(e.Timestamp.Format("2006-01-02 15:04"))
			mk := marginStyle.Render(e.MarginKey)
			blt := bulletStyle.Render(e.Bullet)
			cnt := noteStyle.Render(e.Content)
			id := idStyle.Render(e.ID)

			prio := ""
			if e.Priority {
				prio = priorityStyle.Render("!")
			}

			fmt.Printf("%s | %s | %s | %s %s %s\n", id, ts, mk, blt, cnt, prio)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&timeFilter, "time", "t", "", "Time filter (3d, 7d, today)")
	listCmd.Flags().StringVarP(&filterMargin, "filter", "f", "", "Filter by margin key (e.g. work)")
	listCmd.Flags().BoolVarP(&filterPriority, "priority", "p", false, "Filter by priority")
	rootCmd.AddCommand(listCmd)
}
