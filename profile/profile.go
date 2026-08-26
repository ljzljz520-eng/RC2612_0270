package profile

import (
	"errors"
	"strings"
	"time"
)

type Profile struct {
	ID, Name, Title, Institution, Bio, NotificationAddress string
	Tags                                                   []string
	UpdatedAt                                              time.Time
}

func New(id, name string) Profile { return Profile{ID: id, Name: name, UpdatedAt: time.Now().UTC()} }
func (p Profile) Validate() error {
	if strings.TrimSpace(p.ID) == "" || strings.TrimSpace(p.Name) == "" {
		return errors.New("profile id and name required")
	}
	if len(p.Name) > 160 {
		return errors.New("name too long")
	}
	return nil
}
func (p *Profile) Normalize() {
	p.Name = strings.TrimSpace(p.Name)
	p.Title = strings.TrimSpace(p.Title)
	p.Institution = strings.TrimSpace(p.Institution)
	p.Bio = strings.TrimSpace(p.Bio)
	for i := range p.Tags {
		p.Tags[i] = strings.TrimSpace(p.Tags[i])
	}
	p.UpdatedAt = time.Now().UTC()
}
func (p Profile) PublicSummary() map[string]any {
	return map[string]any{"id": p.ID, "name": p.Name, "title": p.Title, "institution": p.Institution, "tags": append([]string(nil), p.Tags...)}
}
func (p Profile) HasTag(tag string) bool {
	for _, t := range p.Tags {
		if strings.EqualFold(t, tag) {
			return true
		}
	}
	return false
}
func Merge(base, patch Profile) Profile {
	if patch.Name != "" {
		base.Name = patch.Name
	}
	if patch.Title != "" {
		base.Title = patch.Title
	}
	if patch.Institution != "" {
		base.Institution = patch.Institution
	}
	if patch.Bio != "" {
		base.Bio = patch.Bio
	}
	if patch.Tags != nil {
		base.Tags = patch.Tags
	}
	if patch.NotificationAddress != "" {
		base.NotificationAddress = patch.NotificationAddress
	}
	base.Normalize()
	return base
}
