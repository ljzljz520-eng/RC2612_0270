package store

import (
	"frontend_go/records"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "reopen.db")
	d, e := Open(path)
	if e != nil {
		t.Fatal(e)
	}
	r := records.New("persist", "p", "Persistent Study")
	r.Kind = "paper"
	if e = d.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if e = d.Close(); e != nil {
		t.Fatal(e)
	}
	d, e = Open(path)
	if e != nil {
		t.Fatal(e)
	}
	defer d.Close()
	got, e := d.LoadRecord("persist")
	if e != nil || got.Title != "Persistent Study" {
		t.Fatalf("reopen: %v %#v", e, got)
	}
}
