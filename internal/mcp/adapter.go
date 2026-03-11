package mcp

import (
	"context"
	"rapide/internal/model"
	"rapide/internal/storage"
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
		if e.MarginKey == "AGENT" && strings.Contains(strings.ToLower(e.Content), q) {
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
		if all[i].MarginKey == "AGENT" {
			agentEntries = append(agentEntries, all[i])
			if len(agentEntries) >= limit {
				break
			}
		}
	}
	return agentEntries, nil
}
