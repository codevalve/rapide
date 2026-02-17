package storage

import (
	"os"
	"path/filepath"
	"rapide/internal/model"
	"testing"
)

func TestStorage(t *testing.T) {
	// Setup temp storage
	tmpDir, err := os.MkdirTemp("", "rapide-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	dbPath := filepath.Join(tmpDir, "entries.jsonl")
	s := &Storage{FilePath: dbPath}

	// Test Append
	entry := model.Entry{
		Content: "Test entry",
		Bullet:  "•",
	}
	id, err := s.Append(entry)
	if err != nil {
		t.Fatalf("Append failed: %v", err)
	}
	if id == "" {
		t.Error("expected non-empty ID")
	}

	// Test List
	entries, err := s.List()
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(entries) != 1 {
		t.Errorf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].ID != id {
		t.Errorf("expected ID %s, got %s", id, entries[0].ID)
	}

	// Test Update
	entries[0].Content = "Updated content"
	err = s.Update(id, entries[0])
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	// Verify Update
	entries, _ = s.List()
	if entries[0].Content != "Updated content" {
		t.Errorf("expected content 'Updated content', got '%s'", entries[0].Content)
	}
}
