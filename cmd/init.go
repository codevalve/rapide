package cmd

import (
	"bufio"
	"fmt"
	"os"
	"rapide/internal"
	"rapide/internal/model"
	"rapide/internal/storage"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Rapide with a guided setup wizard",
	Long:  "Run the Rapide setup wizard to configure your journal and seed it with example entries.",
	Run: func(cmd *cobra.Command, args []string) {
		runInitWizard()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// promptStyle renders a bolded cyan prompt label.
var promptStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("6"))

// dimStyle renders muted hint text.
var dimStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("8"))

// accentStyle renders highlighted user input / confirmations.
var accentStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("208"))

// bannerStyle for the header block.
var bannerStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("6")).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("6")).
	Padding(0, 2)

func ask(reader *bufio.Reader, question, defaultVal string) string {
	hint := ""
	if defaultVal != "" {
		hint = dimStyle.Render(fmt.Sprintf(" [%s]", defaultVal))
	}
	fmt.Printf("%s%s: ", promptStyle.Render(question), hint)
	raw, _ := reader.ReadString('\n')
	val := strings.TrimSpace(raw)
	if val == "" {
		return defaultVal
	}
	return val
}

func confirm(reader *bufio.Reader, question string) bool {
	fmt.Printf("%s %s ", promptStyle.Render(question), dimStyle.Render("[y/N]"))
	raw, _ := reader.ReadString('\n')
	ans := strings.TrimSpace(strings.ToLower(raw))
	return ans == "y" || ans == "yes"
}

func runInitWizard() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println(bannerStyle.Render("  🗿 Welcome to RAPIDE  "))
	fmt.Println()
	fmt.Println(dimStyle.Render("  Let's get your journal set up. Press Enter to accept defaults."))
	fmt.Println()

	// Check existing journal
	s, err := storage.NewStorage()
	if err != nil {
		fmt.Printf("%s Could not access storage: %v\n", accentStyle.Render("✗"), err)
		os.Exit(1)
	}

	existing, _ := s.List()
	if len(existing) > 0 {
		fmt.Printf("  %s Journal already contains %s entries.\n",
			accentStyle.Render("!"),
			accentStyle.Render(fmt.Sprintf("%d", len(existing))))
		fmt.Println()
		if !confirm(reader, "Seed with example entries anyway?") {
			fmt.Println()
			fmt.Println(dimStyle.Render("  Init cancelled. Your journal is unchanged."))
			fmt.Println()
			return
		}
		fmt.Println()
	}

	// --- Step 1: Name ---
	fmt.Printf("  %s\n", dimStyle.Render("Step 1 of 3: About You"))
	name := ask(reader, "  Your name", "Journaler")
	fmt.Println()

	// --- Step 2: Primary margin keys ---
	fmt.Printf("  %s\n", dimStyle.Render("Step 2 of 3: Your Collections"))
	fmt.Printf("  %s\n", dimStyle.Render("Collections are margin keys that organize your entries (e.g. work, health, ideas)."))
	rawKeys := ask(reader, "  Collections (comma-separated)", "work,health,ideas")
	marginKeys := []string{}
	for _, k := range strings.Split(rawKeys, ",") {
		k = strings.TrimSpace(k)
		if k != "" {
			marginKeys = append(marginKeys, k)
		}
	}
	if len(marginKeys) == 0 {
		marginKeys = []string{"work", "health", "ideas"}
	}
	fmt.Println()

	// --- Step 3: Git sync ---
	fmt.Printf("  %s\n", dimStyle.Render("Step 3 of 3: Git Sync (optional)"))
	fmt.Printf("  %s\n", dimStyle.Render("Link a private Git repo to sync your journal across machines."))
	remoteURL := ask(reader, "  Remote Git URL (leave blank to skip)", "")
	autoSync := false
	if remoteURL != "" {
		autoSync = confirm(reader, "  Enable auto-sync on every write?")
	}
	fmt.Println()

	// Save config if a remote was provided
	if remoteURL != "" {
		cfg := &storage.Config{
			RemoteURL:    remoteURL,
			AutoSync:     autoSync,
			AutoHideDays: 14,
		}
		if err := storage.SaveConfig(cfg); err != nil {
			fmt.Printf("  %s Could not save config: %v\n", accentStyle.Render("✗"), err)
		} else {
			fmt.Printf("  %s Git config saved.\n", successStyle.Render("✓"))
		}
	}

	// --- Seed entries ---
	now := time.Now()
	entries := buildSeedEntries(name, marginKeys, now)

	pinned := false
	for i, e := range entries {
		id, err := s.Append(e)
		if err != nil {
			fmt.Printf("  %s Failed to add entry: %v\n", accentStyle.Render("✗"), err)
			continue
		}
		// Pin the welcome note (first entry)
		if !pinned {
			s.TogglePin(id)
			pinned = true
		}
		_ = i
	}

	fmt.Println()
	fmt.Printf("  %s %s entries added to your journal.\n",
		successStyle.Render("✓"),
		accentStyle.Render(fmt.Sprintf("%d", len(entries))))
	fmt.Println()
	fmt.Printf("  %s  Run %s to open the interactive journal.\n",
		dimStyle.Render("→"),
		accentStyle.Render("rapide tui"))
	fmt.Printf("  %s  Run %s to log your first entry.\n",
		dimStyle.Render("→"),
		accentStyle.Render("rapide • My first task"))
	fmt.Println()
}

// buildSeedEntries creates curated starter entries showcasing all bullet types.
func buildSeedEntries(name string, marginKeys []string, base time.Time) []model.Entry {
	entries := []model.Entry{}

	// Welcome note — always added, pinned later
	entries = append(entries, makeEntry("", "•",
		fmt.Sprintf("Welcome to Rapide, %s! Press ? in the TUI for help.", name),
		base.Add(-1*time.Minute), false))

	// Syntax guide note
	entries = append(entries, makeEntry("", "-",
		"Syntax: [collection | ] [bullet] content [!]  — use ! for priority",
		base.Add(-2*time.Minute), false))

	// Per-collection examples
	for i, mk := range marginKeys {
		offset := time.Duration(-(i+1)*10) * time.Minute
		switch i % 3 {
		case 0: // task-heavy
			entries = append(entries, makeEntry(mk, "•", "Example task — mark done with 'd' in TUI", base.Add(offset), false))
			entries = append(entries, makeEntry(mk, "•", "Priority task example!", base.Add(offset-1*time.Minute), true))
			entries = append(entries, makeEntry(mk, "O", "Example event on the calendar", base.Add(offset-2*time.Minute), false))
		case 1: // notes
			entries = append(entries, makeEntry(mk, "-", "Example note — observations go here", base.Add(offset), false))
			entries = append(entries, makeEntry(mk, "AI", "Example action item — assigned to self", base.Add(offset-1*time.Minute), false))
		case 2: // mixed
			entries = append(entries, makeEntry(mk, "•", "Something to work on soon", base.Add(offset), false))
			entries = append(entries, makeEntry(mk, "-", "A thought to capture", base.Add(offset-1*time.Minute), false))
		}
	}

	return entries
}

func makeEntry(marginKey, bullet, content string, ts time.Time, priority bool) model.Entry {
	parsed := internal.ParseEntry([]string{bullet + " " + content})
	parsed.MarginKey = marginKey
	parsed.Timestamp = ts
	parsed.Priority = priority
	return parsed
}
