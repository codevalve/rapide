package model

import "time"

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	MarginKey string    `json:"margin_key,omitempty"`
	Bullet    string    `json:"bullet"`
	Content   string    `json:"content"`
	Priority  bool      `json:"priority"`
}
