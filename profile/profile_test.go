package profile

import "testing"

func TestProfileValidation(t *testing.T) {
	p := New("p1", " Ada ")
	p.Normalize()
	if p.Name != "Ada" {
		t.Fatal(p.Name)
	}
	if p.Validate() != nil {
		t.Fatal("valid profile rejected")
	}
}
