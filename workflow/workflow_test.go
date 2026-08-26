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

func setup(t *testing.T) (*Service, func()) {
	d, e := store.Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	return New(d, notify.New(func(context.Context, string) error { return nil })), func() { d.Close() }
}
func TestWorkflowOne(t *testing.T) {
	s, done := setup(t)
	defer done()
	if e := s.RegisterProfile(profile.New("p", "Researcher")); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, done := setup(t)
	defer done()
	r := records.New("r", "p", "Study")
	r.Kind = "paper"
	r.SetAbstract("A long enough abstract to pass publishing validation.")
	if e := s.SubmitRecord(r); e != nil {
		t.Fatal(e)
	}
	if e := s.ReviewRecord("r", "mentor"); e != nil {
		t.Fatal(e)
	}
	if e := s.ArchiveRecord("r", "mentor"); e == nil {
		t.Fatal("premature archive")
	}
}
func TestWorkflowThree(t *testing.T) {
	s, done := setup(t)
	defer done()
	if e := s.PublishRecord(context.Background(), "missing", "a"); e == nil {
		t.Fatal("missing record")
	}
}
