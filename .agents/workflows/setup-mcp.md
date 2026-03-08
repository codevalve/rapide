---
description: Blueprint for implementing Model Context Protocol (MCP) into Rapide.
---

# MCP Integration Workflow (v3.0)

This workflow outlines the steps to turn Rapide into a first-class MCP server.

## 1. Initial Research
- [ ] View MCP specification at [modelcontextprotocol.io](https://modelcontextprotocol.io)
- [ ] Check existing Go MCP SDK implementations (e.g., `github.com/mark3labs/mcp-go`)

## 2. Command Scaffolding
- [ ] Add a new subcommand: `rapide mcp start`
- [ ] This command should run a long-lived process listening on `stdio` (default for MCP).

## 3. Resource Mapping
- [ ] Define entry points for journal reading.
- [ ] Map `list` and `search` to MCP "Resources".
- [ ] Example URI: `rapide://all-entries`, `rapide://today`.

## 4. Tool Mapping
- [ ] Expose `AddEntry` as an MCP tool.
- [ ] Expose `ToggleDone` as an MCP tool.

## 5. Security & Privacy
- [ ] Ensure the MCP server only has access to the `~/.rapide/` directory.
- [ ] Implement a toggle in `config.json` to enable/disable MCP mode.
