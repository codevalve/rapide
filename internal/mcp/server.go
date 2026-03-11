package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Server struct {
	adapter JournalAdapter
	stdin   io.Reader
	stdout  io.Writer
}

func NewServer(adapter JournalAdapter) *Server {
	return &Server{
		adapter: adapter,
		stdin:   os.Stdin,
		stdout:  os.Stdout,
	}
}

func (s *Server) Start(ctx context.Context) error {
	scanner := bufio.NewScanner(s.stdin)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			line := scanner.Bytes()
			if len(line) == 0 {
				continue
			}

			var req Request
			if err := json.Unmarshal(line, &req); err != nil {
				s.sendError(nil, -32700, "Parse error", nil)
				continue
			}

			s.handleRequest(ctx, req)
		}
	}
	return scanner.Err()
}

func (s *Server) handleRequest(ctx context.Context, req Request) {
	var result interface{}
	var err *Error

	switch req.Method {
	case "initialize":
		result = s.handleInitialize(req)
	case "tools/list":
		result = s.handleListTools()
	case "tools/call":
		result, err = s.handleCallTool(ctx, req.Params)
	case "notifications/initialized":
		return // Ignore
	default:
		err = &Error{Code: -32601, Message: "Method not found"}
	}

	if req.ID != nil {
		if err != nil {
			s.sendError(req.ID, err.Code, err.Message, err.Data)
		} else {
			s.sendResponse(req.ID, result)
		}
	}
}

func (s *Server) handleInitialize(req Request) InitializeResult {
	res := InitializeResult{
		ProtocolVersion: "2024-11-05",
		ServerInfo: struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{
			Name:    "rapide-mcp",
			Version: "1.0.0",
		},
	}
	res.Capabilities.Tools = struct{}{}
	return res
}

func (s *Server) handleListTools() ListToolsResult {
	return ListToolsResult{
		Tools: []Tool{
			{
				Name:        "add_entry",
				Description: "Add a new rapid-log entry to the journal.",
				InputSchema: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"content": map[string]interface{}{
							"type":        "string",
							"description": "The content of the entry.",
						},
					},
					"required": []string{"content"},
				},
			},
			{
				Name:        "search_agent_entries",
				Description: "Search entries in the AGENT collection.",
				InputSchema: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"query": map[string]interface{}{
							"type":        "string",
							"description": "The search term.",
						},
					},
					"required": []string{"query"},
				},
			},
			{
				Name:        "list_recent_agent_entries",
				Description: "List recent AGENT entries.",
				InputSchema: map[string]interface{}{
					"type": "object",
					"properties": map[string]interface{}{
						"limit": map[string]interface{}{
							"type":        "integer",
							"description": "Number of entries to return.",
							"default":     10,
						},
					},
				},
			},
		},
	}
}

func (s *Server) handleCallTool(ctx context.Context, params json.RawMessage) (interface{}, *Error) {
	var call CallToolParams
	if err := json.Unmarshal(params, &call); err != nil {
		return nil, &Error{Code: -32602, Message: "Invalid params"}
	}

	switch call.Name {
	case "add_entry":
		return s.toolAddEntry(ctx, call.Arguments)
	case "search_agent_entries":
		return s.toolSearchEntries(ctx, call.Arguments)
	case "list_recent_agent_entries":
		return s.toolListRecent(ctx, call.Arguments)
	default:
		return nil, &Error{Code: -32601, Message: "Tool not found"}
	}
}

func (s *Server) sendResponse(id interface{}, result interface{}) {
	res := Response{
		JSONRPC: "2.0",
		ID:      id,
	}
	if result != nil {
		data, _ := json.Marshal(result)
		res.Result = data
	}
	s.writeLine(res)
}

func (s *Server) sendError(id interface{}, code int, message string, data interface{}) {
	res := Response{
		JSONRPC: "2.0",
		ID:      id,
		Error: &Error{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}
	s.writeLine(res)
}

func (s *Server) writeLine(v interface{}) {
	data, _ := json.Marshal(v)
	fmt.Fprintf(s.stdout, "%s\n", data)
}
