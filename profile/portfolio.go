package profile

import "sort"

type Portfolio struct {
	ProfileID    string
	Publications []string
	Degrees      []string
	Interests    map[string]int
}

func (p Portfolio) PublicationCount() int { return len(p.Publications) }
func (p *Portfolio) AddPublication(title string) {
	if title == "" {
		return
	}
	for _, x := range p.Publications {
		if x == title {
			return
		}
	}
	p.Publications = append(p.Publications, title)
}
func (p *Portfolio) RemovePublication(title string) {
	out := p.Publications[:0]
	for _, x := range p.Publications {
		if x != title {
			out = append(out, x)
		}
	}
	p.Publications = out
}
func (p Portfolio) RankedInterests() []string {
	out := make([]string, 0, len(p.Interests))
	for k := range p.Interests {
		out = append(out, k)
	}
	sort.Slice(out, func(i, j int) bool { return p.Interests[out[i]] > p.Interests[out[j]] })
	return out
}
func (p *Portfolio) RecordInterest(name string, weight int) {
	if weight < 1 {
		weight = 1
	}
	p.Interests[name] += weight
}
