package notify

import (
	"context"
	"testing"
)

func TestNotifierHistory(t *testing.T) {
	n := New(func(context.Context, string) error { return nil })
	if e := n.Send(context.Background(), "x"); e != nil {
		t.Fatal(e)
	}
	if len(n.History()) != 1 {
		t.Fatal("history")
	}
}
