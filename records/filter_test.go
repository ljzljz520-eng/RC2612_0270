package records

import (
	"testing"
	"time"
)

func TestFilteringAndPaging(t *testing.T) {
	rs := []Record{{ID: "1", Title: "Go", Abstract: "systems", Status: "published", CreatedAt: time.Now()}, {ID: "2", Title: "Rust", Status: "draft"}}
	if len(Paginate(SortNewest(rs), 0, 1)) != 1 {
		t.Fatal("paging")
	}
	if !Match(rs[0], Filter{Query: "go"}) {
		t.Fatal("query")
	}
}
