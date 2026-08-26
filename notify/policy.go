package notify

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
)

type Policy struct {
	Enabled   bool
	Prefix    string
	MaxLength int
}

func (p Policy) Validate() error {
	if p.MaxLength < 0 {
		return errors.New("negative maximum")
	}
	return nil
}
func (p Policy) Format(message string) string {
	message = strings.TrimSpace(message)
	if p.Prefix != "" {
		message = p.Prefix + ": " + message
	}
	if p.MaxLength > 0 && len(message) > p.MaxLength {
		message = message[:p.MaxLength]
	}
	return message
}

type Counter struct{ value atomic.Int64 }

func (c *Counter) Add()         { c.value.Add(1) }
func (c *Counter) Value() int64 { return c.value.Load() }
func Deliver(ctx context.Context, n *Notifier, p Policy, msg string) (bool, error) {
	if !p.Enabled {
		return false, nil
	}
	if err := p.Validate(); err != nil {
		return false, err
	}
	if err := n.Send(ctx, p.Format(msg)); err != nil {
		return false, err
	}
	return true, nil
}
