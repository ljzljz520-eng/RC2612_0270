package store

import (
	"frontend_go/profile"
	"frontend_go/records"
	"path/filepath"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	d, e := Open(filepath.Join(t.TempDir(), "x.db"))
	if e != nil {
		t.Fatal(e)
	}
	p := profile.New("p", "Name")
	if e = d.SaveProfile(p); e != nil {
		t.Fatal(e)
	}
	r := records.New("r", "p", "Title")
	r.Kind = "paper"
	if e = d.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if _, e = d.LoadRecord("r"); e != nil {
		t.Fatal(e)
	}
	d.Close()
}
