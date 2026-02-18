package storage

import (
	"os"
	"path/filepath"
	"rapide/internal/model"
	"testing"
	"time"
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

	// Test Delete
	err = s.Delete(id)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	entries, _ = s.List()
	if len(entries) != 0 {
		t.Errorf("expected 0 entries, got %d", len(entries))
	}
}

func TestTrimAndArchive(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "rapide-trim-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	dbPath := filepath.Join(tmpDir, "entries.jsonl")
	s := &Storage{FilePath: dbPath}

	now := time.Now()
	oldDate := now.AddDate(0, 0, -10)
	newDate := now.AddDate(0, 0, 10)

	// Add old and new entries
	s.Append(model.Entry{Content: "old item", Timestamp: oldDate})
	s.Append(model.Entry{Content: "new item", Timestamp: newDate})

	// Test Trim
	cutoff := now
	count, err := s.TrimBefore(cutoff)
	if err != nil {
		t.Fatalf("TrimBefore failed: %v", err)
	}
	if count != 1 {
		t.Errorf("expected 1 item trimmed, got %d", count)
	}

	entries, _ := s.List()
	if len(entries) != 1 {
		t.Errorf("expected 1 item remaining, got %d", len(entries))
	}
	if entries[0].Content != "new item" {
		t.Errorf("expected 'new item' to remain, got '%s'", entries[0].Content)
	}

	// Reset for Archive test
	os.Remove(dbPath)
	s.Append(model.Entry{Content: "old archive item", Timestamp: oldDate})
	s.Append(model.Entry{Content: "new item", Timestamp: newDate})

	// Test Archive
	count, filename, err := s.ArchiveBefore(cutoff)
	if err != nil {
		t.Fatalf("ArchiveBefore failed: %v", err)
	}
	if count != 1 {
		t.Errorf("expected 1 item archived, got %d", count)
	}
	if filename == "" {
		t.Error("expected non-empty archive filename")
	}

	// Verify archive file exists
	archivePath := filepath.Join(tmpDir, filename)
	if _, err := os.Stat(archivePath); os.IsNotExist(err) {
		t.Errorf("expected archive file %s to exist", archivePath)
	}

	// Verify entries remaining
	entries, _ = s.List()
	if len(entries) != 1 {
		t.Errorf("expected 1 item remaining in main log, got %d", len(entries))
	}
}
