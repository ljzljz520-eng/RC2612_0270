package workflow

import (
	"context"
	"frontend_go/notify"
	"frontend_go/profile"
	"frontend_go/records"
	"frontend_go/store"
	"path/filepath"
	"testing"
)

func TestBusinessChain12(t *testing.T) {
	d, e := store.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer d.Close()
	s := New(d, notify.New(nil))
	p := profile.New("p", "No Notifications")
	if e = s.RegisterProfile(p); e != nil {
		t.Fatal(e)
	}
	r := records.New("r", "p", "Result")
	r.Kind = "paper"
	r.SetAbstract("This abstract is deliberately long enough for release.")
	if e = s.SubmitRecord(r); e != nil {
		t.Fatal(e)
	}
	if e = s.ReviewRecord("r", "mentor"); e != nil {
		t.Fatal(e)
	}
	if e = s.PublishRecord(context.Background(), "r", "mentor"); e != nil {
		t.Fatal(e)
	}
}
