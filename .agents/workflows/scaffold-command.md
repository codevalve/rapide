---
description: How to add a new CLI command to Rapide.
---

# Scaffold Command Workflow

This workflow describes how to add a new subcommand to the Rapide CLI using Cobra.

## 1. Create Command File
New commands should live in the `cmd/` directory.
- [ ] Create a new file `cmd/<command_name>.go`.
- [ ] Set the package to `cmd`.

## 2. Define Command Structure
// turbo
- [ ] Add the following boilerplate:
```go
package cmd

import (
	"github.com/spf13/cobra"
)

var <name>Cmd = &cobra.Command{
	Use:   "<name>",
	Short: "A brief description of the command",
	Long:  `A longer description of the command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
	},
}

func init() {
	rootCmd.AddCommand(<name>Cmd)
}
```

## 3. Implement Logic
- [ ] Use `internal/storage` for data access.
- [ ] Use `internal/tui/styles.go` for any themed output if necessary, though CLI-specific styles can also be placed in `cmd/root.go`.

## 4. Register Subcommand
The `init()` function automatically registers the command with `rootCmd`.

## 5. Verification
// turbo
- [ ] Run `go build -o rapide`
- [ ] Run `./rapide <name> --help` to verify registration.
