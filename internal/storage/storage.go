package storage

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"rapide/internal/model"
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

func (s *Storage) Append(entry model.Entry) error {
	f, err := os.OpenFile(s.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	_, err = f.Write(append(data, '\n'))
	return err
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
		entries = append(entries, entry)
	}
	return entries, scanner.Err()
}
