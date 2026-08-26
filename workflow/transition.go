package workflow

import (
	"errors"
	"fmt"
	"frontend_go/records"
)

var transitions = map[string]map[string]bool{"draft": {"submitted": true}, "submitted": {"reviewed": true}, "reviewed": {"published": true}, "published": {"archived": true}}

func CanTransition(from, to string) bool { return transitions[from][to] }
func Transition(r *records.Record, to string) error {
	if r == nil {
		return errors.New("record required")
	}
	if !CanTransition(r.Status, to) {
		return fmt.Errorf("invalid transition %s to %s", r.Status, to)
	}
	r.Status = to
	r.Version++
	return nil
}
func Pending(status string) bool { return status == "submitted" || status == "reviewed" }
func WorkflowStages() []string {
	return []string{"draft", "submitted", "reviewed", "published", "archived"}
}
