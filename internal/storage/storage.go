package storage

import (
	"bufio"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"rapide/internal/model"
	"time"
)

type Storage struct {
	FilePath string
}

func NewStorage() (*Storage, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dir := filepath.Join(home, ".rapide")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	return &Storage{FilePath: filepath.Join(dir, "entries.jsonl")}, nil
}

func (s *Storage) Append(entry model.Entry) (string, error) {
	if entry.ID == "" {
		entry.ID = generateID(entry)
	}

	f, err := os.OpenFile(s.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data, err := json.Marshal(entry)
	if err != nil {
		return "", err
	}

	_, err = f.Write(append(data, '\n'))
	return entry.ID, err
}

func (s *Storage) List() ([]model.Entry, error) {
	f, err := os.Open(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []model.Entry{}, nil
		}
		return nil, err
	}
	defer f.Close()

	var entries []model.Entry
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var entry model.Entry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			continue // skip malformed lines
		}
		// Migration for old entries without IDs
		if entry.ID == "" {
			entry.ID = generateID(entry)
		}
		entries = append(entries, entry)
	}
	return entries, scanner.Err()
}

func (s *Storage) Update(id string, newEntry model.Entry) error {
	entries, err := s.List()
	if err != nil {
		return err
	}

	found := false
	for i, e := range entries {
		if e.ID == id {
			newEntry.ID = id
			if newEntry.Timestamp.IsZero() {
				newEntry.Timestamp = e.Timestamp
			}
			entries[i] = newEntry
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("entry with ID %s not found", id)
	}

	return s.saveAll(entries)
}

func (s *Storage) Delete(id string) error {
	entries, err := s.List()
	if err != nil {
		return err
	}

	var updated []model.Entry
	found := false
	for _, e := range entries {
		if e.ID == id {
			found = true
			continue
		}
		updated = append(updated, e)
	}

	if !found {
		return fmt.Errorf("entry with ID %s not found", id)
	}

	return s.saveAll(updated)
}

func (s *Storage) saveAll(entries []model.Entry) error {
	f, err := os.OpenFile(s.FilePath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, e := range entries {
		data, err := json.Marshal(e)
		if err != nil {
			return err
		}
		if _, err := f.Write(append(data, '\n')); err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) ArchiveCompleted() (int, string, error) {
	entries, err := s.List()
	if err != nil {
		return 0, "", err
	}

	var toKeep, toArchive []model.Entry
	for _, e := range entries {
		if e.Bullet == "x" || e.Bullet == ">" {
			toArchive = append(toArchive, e)
		} else {
			toKeep = append(toKeep, e)
		}
	}

	if len(toArchive) == 0 {
		return 0, "", nil
	}

	archiveName := fmt.Sprintf("archive_completed_%s.jsonl", time.Now().Format("20060102_150405"))
	archivePath := filepath.Join(filepath.Dir(s.FilePath), archiveName)

	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return 0, "", err
	}
	defer archiveFile.Close()

	for _, e := range toArchive {
		data, err := json.Marshal(e)
		if err != nil {
			return 0, "", err
		}
		if _, err := archiveFile.Write(append(data, '\n')); err != nil {
			return 0, "", err
		}
	}

	if err := s.saveAll(toKeep); err != nil {
		return 0, "", err
	}

	return len(toArchive), archiveName, nil
}

func (s *Storage) ArchiveBefore(cutoff time.Time) (int, string, error) {
	entries, err := s.List()
	if err != nil {
		return 0, "", err
	}

	var toKeep, toArchive []model.Entry
	for _, e := range entries {
		if e.Timestamp.Before(cutoff) {
			toArchive = append(toArchive, e)
		} else {
			toKeep = append(toKeep, e)
		}
	}

	if len(toArchive) == 0 {
		return 0, "", nil
	}

	archiveName := fmt.Sprintf("archive_%s.jsonl", time.Now().Format("20060102_150405"))
	archivePath := filepath.Join(filepath.Dir(s.FilePath), archiveName)

	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return 0, "", err
	}
	defer archiveFile.Close()

	for _, e := range toArchive {
		data, err := json.Marshal(e)
		if err != nil {
			return 0, "", err
		}
		if _, err := archiveFile.Write(append(data, '\n')); err != nil {
			return 0, "", err
		}
	}

	if err := s.saveAll(toKeep); err != nil {
		return 0, "", err
	}

	return len(toArchive), archiveName, nil
}

func (s *Storage) TrimBefore(cutoff time.Time) (int, error) {
	entries, err := s.List()
	if err != nil {
		return 0, err
	}

	var toKeep []model.Entry
	count := 0
	for _, e := range entries {
		if e.Timestamp.Before(cutoff) {
			count++
			continue
		}
		toKeep = append(toKeep, e)
	}

	if count == 0 {
		return 0, nil
	}

	return count, s.saveAll(toKeep)
}

func generateID(e model.Entry) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%d%s%s", e.Timestamp.UnixNano(), e.Bullet, e.Content)))
	return fmt.Sprintf("%x", h.Sum(nil))[:4]
}
