package store

import (
	"encoding/json"
	"errors"
	"frontend_go/archive"
	"frontend_go/profile"
	"frontend_go/records"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var buckets = [][]byte{[]byte("profiles"), []byte("records"), []byte("events"), []byte("audits")}

type DB struct {
	raw *bbolt.DB
	mu  sync.RWMutex
}

func Open(path string) (*DB, error) {
	if path == "" {
		return nil, errors.New("database path required")
	}
	raw, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	d := &DB{raw: raw}
	err = raw.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		raw.Close()
		return nil, err
	}
	return d, nil
}
func (d *DB) Close() error { d.mu.Lock(); defer d.mu.Unlock(); return d.raw.Close() }
func put(tx *bbolt.Tx, b, key string, v any) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket([]byte(b)).Put([]byte(key), data)
}
func get(tx *bbolt.Tx, b, key string, v any) error {
	data := tx.Bucket([]byte(b)).Get([]byte(key))
	if data == nil {
		return errors.New("not found")
	}
	return json.Unmarshal(data, v)
}
func (d *DB) SaveProfile(p profile.Profile) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return put(tx, "profiles", p.ID, p) })
}
func (d *DB) LoadProfile(id string) (profile.Profile, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var p profile.Profile
	e := d.raw.View(func(tx *bbolt.Tx) error { return get(tx, "profiles", id, &p) })
	return p, e
}
func (d *DB) SaveRecord(r records.Record) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return put(tx, "records", r.ID, r) })
}
func (d *DB) LoadRecord(id string) (records.Record, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	var r records.Record
	e := d.raw.View(func(tx *bbolt.Tx) error { return get(tx, "records", id, &r) })
	return r, e
}
func (d *DB) SaveEvent(e archive.Event) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return put(tx, "events", e.ID, e) })
}
func (d *DB) SaveAudit(a archive.Audit) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return put(tx, "audits", a.ID, a) })
}
func all[T any](d *DB, b []byte) ([]T, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	out := []T{}
	e := d.raw.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(b).ForEach(func(_, v []byte) error {
			var x T
			if err := json.Unmarshal(v, &x); err != nil {
				return err
			}
			out = append(out, x)
			return nil
		})
	})
	return out, e
}
func (d *DB) AllRecords() ([]records.Record, error) { return all[records.Record](d, []byte("records")) }
func (d *DB) AllEvents() ([]archive.Event, error)   { return all[archive.Event](d, []byte("events")) }
