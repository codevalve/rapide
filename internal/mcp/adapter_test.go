package mcp

import (
	"context"
	"os"
	"path/filepath"
	"rapide/internal/storage"
	"testing"
)

func setupTestStorage(t *testing.T) (*storage.Storage, func()) {
	tmpDir, err := os.MkdirTemp("", "rapide-mcp-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	
	dbPath := filepath.Join(tmpDir, "entries.jsonl")
	s := &storage.Storage{FilePath: dbPath}
	
	return s, func() {
		os.RemoveAll(tmpDir)
	}
}

func TestAdapter_AddAgentEntry(t *testing.T) {
	s, cleanup := setupTestStorage(t)
	defer cleanup()
	
	adapter := NewJournalAdapter(s)
	content := "test mcp entry"
	
	entry, err := adapter.AddAgentEntry(context.Background(), content)
	if err != nil {
		t.Fatalf("AddAgentEntry failed: %v", err)
	}
	
	if entry.Content != content {
		t.Errorf("expected content %q, got %q", content, entry.Content)
	}
	if entry.MarginKey != "AGENT" {
		t.Errorf("expected margin key AGENT, got %q", entry.MarginKey)
	}
}

func TestAdapter_SearchAgentEntries(t *testing.T) {
	s, cleanup := setupTestStorage(t)
	defer cleanup()
	
	adapter := NewJournalAdapter(s)
	adapter.AddAgentEntry(context.Background(), "apple pie")
	adapter.AddAgentEntry(context.Background(), "banana split")
	
	results, err := adapter.SearchAgentEntries(context.Background(), "banana")
	if err != nil {
		t.Fatalf("SearchAgentEntries failed: %v", err)
	}
	
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Content != "banana split" {
		t.Errorf("expected 'banana split', got %q", results[0].Content)
	}
}

func TestAdapter_ListRecentAgentEntries(t *testing.T) {
	s, cleanup := setupTestStorage(t)
	defer cleanup()
	
	adapter := NewJournalAdapter(s)
	for i := 0; i < 5; i++ {
		adapter.AddAgentEntry(context.Background(), "entry")
	}
	
	results, err := adapter.ListRecentAgentEntries(context.Background(), 3)
	if err != nil {
		t.Fatalf("ListRecentAgentEntries failed: %v", err)
	}
	
	if len(results) != 3 {
		t.Errorf("expected 3 results, got %d", len(results))
	}
}
