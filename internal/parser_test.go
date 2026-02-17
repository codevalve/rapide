package internal

import (
	"testing"
)

func TestParseEntry(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantBlt  string
		wantCont string
		wantPrio bool
		wantMK   string
	}{
		{
			name:     "Simple note",
			args:     []string{"-", "Hello world"},
			wantBlt:  "-",
			wantCont: "Hello world",
			wantPrio: false,
		},
		{
			name:     "Task with asterisk",
			args:     []string{"*", "Buy milk"},
			wantBlt:  "•",
			wantCont: "Buy milk",
			wantPrio: false,
		},
		{
			name:     "Priority task",
			args:     []string{"•", "Fix bug!"},
			wantBlt:  "•",
			wantCont: "Fix bug",
			wantPrio: true,
		},
		{
			name:     "Combined margin and bullet",
			args:     []string{"work", "|", "-", "Deep work"},
			wantBlt:  "-",
			wantCont: "Deep work",
			wantMK:   "work",
		},
		{
			name:     "Event with margin",
			args:     []string{"personal", "|", "O", "Party"},
			wantBlt:  "O",
			wantCont: "Party",
			wantMK:   "personal",
		},
		{
			name:     "Done task",
			args:     []string{"x", "Finished task"},
			wantBlt:  "x",
			wantCont: "Finished task",
		},
		{
			name:     "Migrated task",
			args:     []string{">", "Moved task"},
			wantBlt:  ">",
			wantCont: "Moved task",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseEntry(tt.args)
			if got.Bullet != tt.wantBlt {
				t.Errorf("ParseEntry() Bullet = %v, want %v", got.Bullet, tt.wantBlt)
			}
			if got.Content != tt.wantCont {
				t.Errorf("ParseEntry() Content = %v, want %v", got.Content, tt.wantCont)
			}
			if got.Priority != tt.wantPrio {
				t.Errorf("ParseEntry() Priority = %v, want %v", got.Priority, tt.wantPrio)
			}
			if got.MarginKey != tt.wantMK {
				t.Errorf("ParseEntry() MarginKey = %v, want %v", got.MarginKey, tt.wantMK)
			}
		})
	}
}
