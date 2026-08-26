package records

import (
	"testing"
	"time"
)

func TestRecordLifecycle(t *testing.T) {
	r := New("r1", "p1", "Thesis")
	r.Kind = "paper"
	r.SetAbstract("A sufficiently long abstract for publication.")
	if e := r.Publish(time.Now()); e != nil {
		t.Fatal(e)
	}
	if e := r.Archive(); e != nil {
		t.Fatal(e)
	}
}
