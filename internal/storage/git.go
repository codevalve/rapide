package storage

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (s *Storage) Sync() error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	if cfg.RemoteURL == "" {
		return fmt.Errorf("no remote repository configured. run 'rapide sync --setup' first")
	}

	dir := filepath.Dir(s.FilePath)

	// Check if git is initialized
	if _, err := os.Stat(filepath.Join(dir, ".git")); os.IsNotExist(err) {
		return fmt.Errorf("git repository not found in %s", dir)
	}

	// 1. Add and commit local changes if any
	// We track entries.jsonl and config.json
	relPath, _ := filepath.Rel(dir, s.FilePath)
	configPath := "config.json"

	files := []string{relPath, configPath}
	var changedFiles []string

	for _, f := range files {
		cmdStatus := exec.Command("git", "status", "--porcelain", f)
		cmdStatus.Dir = dir
		out, _ := cmdStatus.Output()
		if len(out) > 0 {
			changedFiles = append(changedFiles, f)
		}
	}

	if len(changedFiles) > 0 {
		// Stage changes
		args := append([]string{"add"}, changedFiles...)
		cmdAdd := exec.Command("git", args...)
		cmdAdd.Dir = dir
		if err := cmdAdd.Run(); err != nil {
			return fmt.Errorf("git add failed: %w", err)
		}

		// Commit with -m and --no-edit. We also set GIT_EDITOR to true to prevent any interactive prompt.
		cmdCommit := exec.Command("git", "commit", "-a", "-m", "rapide: auto sync update", "--no-edit")
		cmdCommit.Dir = dir
		cmdCommit.Env = append(os.Environ(), "GIT_EDITOR=true")
		if err := cmdCommit.Run(); err != nil {
			return fmt.Errorf("git commit failed: %w", err)
		}
	}

	// 2. Fetch and Rebase
	// Note: fetch will fail if the remote branch doesn't exist yet (first sync)
	// We check if the remote branch exists before pulling
	cmdCheckRemote := exec.Command("git", "ls-remote", "--heads", "origin", "main")
	cmdCheckRemote.Dir = dir
	remoteExists, _ := cmdCheckRemote.Output()

	if len(remoteExists) > 0 {
		cmdPull := exec.Command("git", "pull", "--rebase", "origin", "main")
		cmdPull.Dir = dir
		if err := cmdPull.Run(); err != nil {
			return fmt.Errorf("git pull failed: %w. manual conflict resolution might be needed in %s", err, s.FilePath)
		}
	}

	// 3. Push
	cmdPush := exec.Command("git", "push", "origin", "main")
	cmdPush.Dir = dir
	if err := cmdPush.Run(); err != nil {
		return fmt.Errorf("git push failed: %w", err)
	}

	return nil
}

func (s *Storage) SetupGit(remoteURL string) error {
	dir := filepath.Dir(s.FilePath)

	// 1. Initialize git if needed
	if _, err := os.Stat(filepath.Join(dir, ".git")); os.IsNotExist(err) {
		cmdInit := exec.Command("git", "init")
		cmdInit.Dir = dir
		if err := cmdInit.Run(); err != nil {
			return fmt.Errorf("git init failed: %w", err)
		}

		// Initial branch should be main
		cmdBranch := exec.Command("git", "checkout", "-b", "main")
		cmdBranch.Dir = dir
		_ = cmdBranch.Run()
	}

	// 2. Set remote
	cmdCheck := exec.Command("git", "remote", "get-url", "origin")
	cmdCheck.Dir = dir
	if err := cmdCheck.Run(); err == nil {
		cmdRemote := exec.Command("git", "remote", "set-url", "origin", remoteURL)
		cmdRemote.Dir = dir
		if err := cmdRemote.Run(); err != nil {
			return fmt.Errorf("git remote set-url failed: %w", err)
		}
	} else {
		cmdRemote := exec.Command("git", "remote", "add", "origin", remoteURL)
		cmdRemote.Dir = dir
		if err := cmdRemote.Run(); err != nil {
			return fmt.Errorf("git remote add failed: %w", err)
		}
	}

	// 3. Save to config first so it can be included in initial commit
	cfg, _ := LoadConfig()
	if cfg == nil {
		cfg = &Config{}
	}
	cfg.RemoteURL = remoteURL
	if err := SaveConfig(cfg); err != nil {
		return err
	}

	// 4. Initial commit and push if there are existing entries
	relPath, _ := filepath.Rel(dir, s.FilePath)
	filesAdded := false
	if _, err := os.Stat(s.FilePath); err == nil {
		cmdAdd := exec.Command("git", "add", relPath)
		cmdAdd.Dir = dir
		cmdAdd.Run()
		filesAdded = true
	}
	if _, err := os.Stat(filepath.Join(dir, "config.json")); err == nil {
		cmdAdd := exec.Command("git", "add", "config.json")
		cmdAdd.Dir = dir
		cmdAdd.Run()
		filesAdded = true
	}

	if filesAdded {
		cmdCommit := exec.Command("git", "commit", "-a", "-m", "rapide: initial journal sync", "--no-edit")
		cmdCommit.Dir = dir
		cmdCommit.Env = append(os.Environ(), "GIT_EDITOR=true")
		cmdCommit.Run()

		// We attempt a push but don't fail setup if it fails (e.g. auth issues)
		cmdPush := exec.Command("git", "push", "-u", "origin", "main")
		cmdPush.Dir = dir
		_ = cmdPush.Run()
	}

	return nil
}
