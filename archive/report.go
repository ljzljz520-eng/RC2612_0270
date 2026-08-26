package archive

import (
	"sort"
	"strings"
	"time"
)

type Report struct {
	RecordID    string
	Actions     []string
	First, Last time.Time
	Complete    bool
}

func BuildReport(events []Event, rid string) Report {
	ts := Timeline(events, rid)
	r := Report{RecordID: rid}
	for _, e := range ts {
		r.Actions = append(r.Actions, e.Action)
		if r.First.IsZero() {
			r.First = e.At
		}
		r.Last = e.At
	}
	r.Complete = len(r.Actions) > 0 && r.Actions[len(r.Actions)-1] == "archived"
	return r
}
func (r Report) Duration() time.Duration {
	if r.First.IsZero() || r.Last.IsZero() {
		return 0
	}
	return r.Last.Sub(r.First)
}
func (r Report) Label() string {
	if r.Complete {
		return "complete"
	}
	if len(r.Actions) == 0 {
		return "empty"
	}
	return strings.Join(r.Actions, " -> ")
}
func MergeReports(a, b Report) Report {
	if a.RecordID == "" {
		return b
	}
	if b.RecordID == "" || a.RecordID != b.RecordID {
		return a
	}
	out := a
	out.Actions = append(append([]string(nil), a.Actions...), b.Actions...)
	sort.Strings(out.Actions)
	if out.First.IsZero() || (!b.First.IsZero() && b.First.Before(out.First)) {
		out.First = b.First
	}
	if b.Last.After(out.Last) {
		out.Last = b.Last
	}
	out.Complete = a.Complete || b.Complete
	return out
}
