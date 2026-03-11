package mcp

import (
	"context"
	"encoding/json"
	"fmt"
)

func (s *Server) toolAddEntry(ctx context.Context, args json.RawMessage) (interface{}, *Error) {
	var a AddEntryArgs
	if err := json.Unmarshal(args, &a); err != nil {
		return nil, &Error{Code: -32602, Message: "Invalid arguments"}
	}

	if a.Content == "" {
		return nil, &Error{Code: -32602, Message: "Content is required"}
	}

	entry, err := s.adapter.AddAgentEntry(ctx, a.Content)
	if err != nil {
		return nil, &Error{Code: -32000, Message: fmt.Sprintf("Failed to add entry: %v", err)}
	}

	return CallToolResult{
		Content: []TextContent{
			{
				Type: "text",
				Text: fmt.Sprintf("Added entry with ID %s to AGENT collection", entry.ID),
			},
		},
	}, nil
}

func (s *Server) toolSearchEntries(ctx context.Context, args json.RawMessage) (interface{}, *Error) {
	var a SearchEntriesArgs
	if err := json.Unmarshal(args, &a); err != nil {
		return nil, &Error{Code: -32602, Message: "Invalid arguments"}
	}

	entries, err := s.adapter.SearchAgentEntries(ctx, a.Query)
	if err != nil {
		return nil, &Error{Code: -32000, Message: fmt.Sprintf("Failed to search entries: %v", err)}
	}

	data, _ := json.MarshalIndent(entries, "", "  ")
	return CallToolResult{
		Content: []TextContent{
			{
				Type: "text",
				Text: string(data),
			},
		},
	}, nil
}

func (s *Server) toolListRecent(ctx context.Context, args json.RawMessage) (interface{}, *Error) {
	var a ListRecentArgs
	if err := json.Unmarshal(args, &a); err != nil {
		// Fallback to default if unmarshal fails (e.g. empty args)
		a.Limit = 10
	}
	if a.Limit <= 0 {
		a.Limit = 10
	}
	if a.Limit > 50 {
		a.Limit = 50
	}

	entries, err := s.adapter.ListRecentAgentEntries(ctx, a.Limit)
	if err != nil {
		return nil, &Error{Code: -32000, Message: fmt.Sprintf("Failed to list entries: %v", err)}
	}

	data, _ := json.MarshalIndent(entries, "", "  ")
	return CallToolResult{
		Content: []TextContent{
			{
				Type: "text",
				Text: string(data),
			},
		},
	}, nil
}
