package notify

import (
	"context"
	"fmt"
	"sync"
)

type Callback func(context.Context, string) error
type Notifier struct {
	callback Callback
	mu       sync.Mutex
	sent     []string
}

func New(cb Callback) *Notifier { return &Notifier{callback: cb} }
func NewConsole() *Notifier {
	return New(func(_ context.Context, msg string) error { fmt.Println(msg); return nil })
}
func (n *Notifier) Send(ctx context.Context, msg string) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	err := n.callback(ctx, msg)
	if err == nil {
		n.sent = append(n.sent, msg)
	}
	return err
}
func (n *Notifier) History() []string {
	n.mu.Lock()
	defer n.mu.Unlock()
	return append([]string(nil), n.sent...)
}
