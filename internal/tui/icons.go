package tui

import (
	"strings"
)

var (
	// Default Unicode symbols (Aesthetic)
	IconPin   = "◆"
	IconClock = "◔"
	IconAgent = "◇"

	PrefixClock = IconClock + " "
	PrefixAgent = IconAgent + " "
)

// SetIconTheme initializes the application with a specific icon theme.
// Supported themes: "unicode", "nerdfont", "octicons"
func SetIconTheme(theme string) {
	switch theme {
	case "nerdfont":
		IconPin = "󰐃"
		IconClock = "󰥔"
		IconAgent = "󰚩"
	case "octicons":
		IconPin = ""
		IconClock = ""
		IconAgent = ""
	default: // "unicode"
		IconPin = "◆"
		IconClock = "◔"
		IconAgent = "◇"
	}
	PrefixClock = IconClock + " "
	PrefixAgent = IconAgent + " "
}

// StripIcons removes aesthetic, legacy, or themed icons from a margin key string.
func StripIcons(s string) string {
	s = strings.TrimPrefix(s, PrefixClock)
	s = strings.TrimPrefix(s, PrefixAgent)
	// Fallback for Nerd Fonts
	s = strings.TrimPrefix(s, "󰥔 ")
	s = strings.TrimPrefix(s, "󰚩 ")
	// Fallback for Octicons
	s = strings.TrimPrefix(s, " ")
	s = strings.TrimPrefix(s, " ")
	// Fallback for legacy icons
	s = strings.TrimPrefix(s, "🕒 ")
	s = strings.TrimPrefix(s, "🤖 ")
	return s
}

// GetIcon returns the appropriate indicator icon based on entry state and theme.
func GetIcon(pinned bool, marginKey string) string {
	if pinned {
		return IconPin
	}
	// Detect clock variants
	if strings.HasPrefix(marginKey, PrefixClock) ||
		strings.HasPrefix(marginKey, "🕒 ") ||
		strings.HasPrefix(marginKey, "󰥔 ") ||
		strings.HasPrefix(marginKey, " ") {
		return IconClock
	}
	// Detect agent variants
	if strings.HasPrefix(marginKey, PrefixAgent) ||
		strings.HasPrefix(marginKey, "🤖 ") ||
		strings.HasPrefix(marginKey, "󰚩 ") ||
		strings.HasPrefix(marginKey, " ") {
		return IconAgent
	}
	return ""
}
