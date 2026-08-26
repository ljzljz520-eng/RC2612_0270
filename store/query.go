package store

import (
	"frontend_go/records"
	"go.etcd.io/bbolt"
	"strings"
)

func (d *DB) FindRecords(f records.Filter) ([]records.Record, error) {
	all, e := d.AllRecords()
	if e != nil {
		return nil, e
	}
	out := make([]records.Record, 0)
	for _, r := range all {
		if f.Query != "" && !strings.Contains(r.SearchText(), strings.ToLower(f.Query)) {
			continue
		}
		if f.Status != "" && r.Status != f.Status {
			continue
		}
		if f.Kind != "" && r.Kind != f.Kind {
			continue
		}
		out = append(out, r)
	}
	return records.SortNewest(out), nil
}
func (d *DB) CountByStatus() map[string]int {
	rs, e := d.AllRecords()
	if e != nil {
		return map[string]int{}
	}
	out := map[string]int{}
	for _, r := range rs {
		out[r.Status]++
	}
	return out
}
func (d *DB) DeleteRecord(id string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.raw.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte("records")).Delete([]byte(id)) })
}
