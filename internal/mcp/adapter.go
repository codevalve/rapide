package mcp

import (
	"context"
	"rapide/internal/model"
	"rapide/internal/storage"
	"rapide/internal/tui"
	"strings"
	"time"
)

type JournalAdapter interface {
	AddAgentEntry(ctx context.Context, content string) (*model.Entry, error)
	SearchAgentEntries(ctx context.Context, query string) ([]model.Entry, error)
	ListRecentAgentEntries(ctx context.Context, limit int) ([]model.Entry, error)
}

type storageAdapter struct {
	storage *storage.Storage
}

func NewJournalAdapter(s *storage.Storage) JournalAdapter {
	return &storageAdapter{storage: s}
}

func (a *storageAdapter) AddAgentEntry(ctx context.Context, content string) (*model.Entry, error) {
	entry := model.Entry{
		Timestamp: time.Now(),
		MarginKey: "AGENT",
		Bullet:    "•",
		Content:   content,
		Priority:  false,
		Pinned:    false,
	}

	// Detect if content starts with HH:MM | or just HH:MM
	parts := strings.SplitN(content, " ", 2)
	potentialTime := strings.TrimSuffix(parts[0], "|")
	if len(potentialTime) == 5 && potentialTime[2] == ':' {
		hh := potentialTime[:2]
		mm := potentialTime[3:]
		// Simple validation
		if hh >= "00" && hh <= "23" && mm >= "00" && mm <= "59" {
			entry.MarginKey = tui.PrefixAgent + potentialTime
			if len(parts) > 1 {
				entry.Content = strings.TrimPrefix(parts[1], "| ")
			} else {
				entry.Content = ""
			}
		}
	}

	id, err := a.storage.Append(entry)
	if err != nil {
		return nil, err
	}
	entry.ID = id
	return &entry, nil
}

func (a *storageAdapter) SearchAgentEntries(ctx context.Context, query string) ([]model.Entry, error) {
	all, err := a.storage.List()
	if err != nil {
		return nil, err
	}

	var results []model.Entry
	q := strings.ToLower(query)
	for _, e := range all {
		isAgent := e.MarginKey == "AGENT" || strings.HasPrefix(e.MarginKey, tui.IconAgent) || strings.HasPrefix(e.MarginKey, "🤖")
		if isAgent && strings.Contains(strings.ToLower(e.Content), q) {
			results = append(results, e)
		}
	}
	return results, nil
}

func (a *storageAdapter) ListRecentAgentEntries(ctx context.Context, limit int) ([]model.Entry, error) {
	all, err := a.storage.List()
	if err != nil {
		return nil, err
	}

	var agentEntries []model.Entry
	for i := len(all) - 1; i >= 0; i-- {
		isAgent := all[i].MarginKey == "AGENT" || strings.HasPrefix(all[i].MarginKey, tui.IconAgent) || strings.HasPrefix(all[i].MarginKey, "🤖")
		if isAgent {
			agentEntries = append(agentEntries, all[i])
			if len(agentEntries) >= limit {
				break
			}
		}
	}
	return agentEntries, nil
}
