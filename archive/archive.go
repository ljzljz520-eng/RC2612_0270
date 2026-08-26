package archive

import (
	"errors"
	"frontend_go/records"
	"sort"
	"time"
)

type Event struct {
	ID, RecordID, Action, Actor string
	At                          time.Time
	Detail                      string
}
type Audit struct {
	ID, Entity, EntityID, Operation, Actor string
	At                                     time.Time
	Payload                                string
}

func NewEvent(id, rid, action, actor string) Event {
	return Event{ID: id, RecordID: rid, Action: action, Actor: actor, At: time.Now().UTC()}
}
func NewAudit(id, entity, entityID, op, actor, payload string) Audit {
	return Audit{ID: id, Entity: entity, EntityID: entityID, Operation: op, Actor: actor, Payload: payload, At: time.Now().UTC()}
}
func Timeline(events []Event, rid string) []Event {
	out := []Event{}
	for _, e := range events {
		if e.RecordID == rid {
			out = append(out, e)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func CanArchive(r records.Record) error {
	if r.Status != "published" {
		return errors.New("record must be published")
	}
	return nil
}
func Summarize(events []Event) map[string]int {
	m := map[string]int{}
	for _, e := range events {
		m[e.Action]++
	}
	return m
}
