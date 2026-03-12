---
description: Blueprint for configuring Rapide as an MCP server.
---

# Setup MCP Workflow

This workflow explains how to expose Rapide as a Model Context Protocol (MCP) server for use with AI agents.

## 1. Prerequisites
- [ ] Rapide v3.0.0 or later installed.
- [ ] An MCP-compatible client (like Claude Desktop or Antigravity).

## 2. Server Command
Rapide includes a hidden command to start the MCP server.
- [ ] Verify the command works locally:
      `rapide mcp start`
      (Note: This will wait for stdio input/output).

## 3. Configuration
Add Rapide to your AI agent's configuration.

### For Antigravity/Claude Desktop:
// turbo
- [ ] Add the following entry to your `mcp_config.json`:
```json
{
  "mcpServers": {
    "rapide": {
      "command": "rapide",
      "args": ["mcp", "start"]
    }
  }
}
```

## 4. Exposed Tools
Once connected, the agent will have access to:
- `add_entry`: Log a new item to the `AGENT` collection.
- `search_agent_entries`: Search through items in the `AGENT` collection.
- `list_recent_agent_entries`: Retrieve the last few items logged by agents.

## 5. Verification
- [ ] Ask the agent to "Log a test entry to Rapide".
- [ ] Verify the entry appears in Rapide with the `AGENT` margin key.
