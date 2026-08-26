package records

import (
	"encoding/json"
	"strings"
	"time"
)

type Snapshot struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Keywords   []string  `json:"keywords"`
	CapturedAt time.Time `json:"captured_at"`
}

func (r Record) Snapshot() Snapshot {
	return Snapshot{ID: r.ID, Title: r.Title, Status: r.Status, Keywords: append([]string(nil), r.Keywords...), CapturedAt: time.Now().UTC()}
}
func (r Record) MarshalSummary() ([]byte, error) { return json.Marshal(r.Snapshot()) }
func ParseKeywords(raw string) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, part := range strings.FieldsFunc(raw, func(r rune) bool { return r == ',' || r == ';' || r == '|' }) {
		v := strings.ToLower(strings.TrimSpace(part))
		if v != "" && !seen[v] {
			seen[v] = true
			out = append(out, v)
		}
	}
	return out
}
func (r *Record) ReplaceKeywords(raw string) { r.Keywords = ParseKeywords(raw); r.Version++ }
func (r Record) IsRecent(now time.Time, days int) bool {
	return !r.CreatedAt.Before(now.AddDate(0, 0, -days))
}
