package internal

import (
	"rapide/internal/model"
	"strings"
	"time"
)

func ParseEntry(args []string) model.Entry {
	fullInput := strings.Join(args, " ")
	entry := model.Entry{
		Timestamp: time.Now(),
	}

	// Check for priority suffix
	if strings.HasSuffix(fullInput, "!") {
		entry.Priority = true
		for strings.HasSuffix(fullInput, "!") {
			fullInput = strings.TrimSuffix(fullInput, "!")
		}
		fullInput = strings.TrimSpace(fullInput)
	}

	parts := strings.SplitN(fullInput, "|", 2)
	var contentPart string

	if len(parts) == 2 {
		entry.MarginKey = strings.TrimSpace(parts[0])
		contentPart = strings.TrimSpace(parts[1])
	} else {
		contentPart = strings.TrimSpace(parts[0])
	}

	// Detect bullet types
	// - (note), O (event), •/* (task), AI/A (action item)
	if strings.HasPrefix(contentPart, "- ") {
		entry.Bullet = "-"
		entry.Content = strings.TrimPrefix(contentPart, "- ")
	} else if strings.HasPrefix(contentPart, "O ") {
		entry.Bullet = "O"
		entry.Content = strings.TrimPrefix(contentPart, "O ")
	} else if strings.HasPrefix(contentPart, "x ") {
		entry.Bullet = "x"
		entry.Content = strings.TrimPrefix(contentPart, "x ")
	} else if strings.HasPrefix(contentPart, "> ") {
		entry.Bullet = ">"
		entry.Content = strings.TrimPrefix(contentPart, "> ")
	} else if strings.HasPrefix(contentPart, "< ") {
		entry.Bullet = "<"
		entry.Content = strings.TrimPrefix(contentPart, "< ")
	} else if strings.HasPrefix(contentPart, "• ") {
		entry.Bullet = "•"
		entry.Content = strings.TrimPrefix(contentPart, "• ")
	} else if strings.HasPrefix(contentPart, "* ") {
		entry.Bullet = "•"
		entry.Content = strings.TrimPrefix(contentPart, "* ")
	} else if strings.HasPrefix(contentPart, "AI ") {
		entry.Bullet = "AI"
		entry.Content = strings.TrimPrefix(contentPart, "AI ")
	} else if strings.HasPrefix(contentPart, "A ") {
		entry.Bullet = "AI"
		entry.Content = strings.TrimPrefix(contentPart, "A ")
	} else {
		// Default to task if no bullet identified?
		// Or just treat as note. Let's default to task (•) if no bullet is found,
		// but the PRD says: "AI Schedule 1:1 with Sarah" -> "AI" is the bullet.
		words := strings.SplitN(contentPart, " ", 2)
		if len(words) > 1 {
			switch words[0] {
			case "AI", "A":
				entry.Bullet = "AI"
				entry.Content = words[1]
			case "-", "O", "•", "*", "x", ">", "<":
				entry.Bullet = words[0]
				if entry.Bullet == "*" {
					entry.Bullet = "•"
				}
				entry.Content = words[1]
			default:
				entry.Bullet = "•" // Default bullet per rapid logging convention often tasks
				entry.Content = contentPart
			}
		} else {
			entry.Bullet = "•"
			entry.Content = contentPart
		}
	}

	return entry
}
