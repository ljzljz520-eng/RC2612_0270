package archive

import (
	"frontend_go/records"
	"testing"
)

func TestTimeline(t *testing.T) {
	e := []Event{NewEvent("1", "r", "published", "a"), NewEvent("2", "x", "published", "a")}
	if len(Timeline(e, "r")) != 1 {
		t.Fatal("timeline")
	}
	if CanArchive(records.Record{Status: "draft"}) == nil {
		t.Fatal("draft archive")
	}
}
