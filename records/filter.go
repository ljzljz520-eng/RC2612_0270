package records

import "sort"

type Filter struct {
	Query, Status, Kind string
	Limit               int
}

func Match(r Record, f Filter) bool {
	if f.Status != "" && r.Status != f.Status {
		return false
	}
	if f.Kind != "" && r.Kind != f.Kind {
		return false
	}
	if f.Query != "" {
		q := f.Query
		if !contains(r.SearchText(), q) {
			return false
		}
	}
	return true
}
func contains(h, n string) bool {
	for i := 0; i+len(n) <= len(h); i++ {
		if h[i:i+len(n)] == n {
			return true
		}
	}
	return n == ""
}
func SortNewest(rs []Record) []Record {
	out := append([]Record(nil), rs...)
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.After(out[j].CreatedAt) })
	return out
}
func Paginate(rs []Record, offset, limit int) []Record {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 25
	}
	if offset >= len(rs) {
		return []Record{}
	}
	end := offset + limit
	if end > len(rs) {
		end = len(rs)
	}
	return rs[offset:end]
}
