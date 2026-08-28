package workflow

import (
	"context"
	"fmt"
	"frontend_go/archive"
	"frontend_go/notify"
	"frontend_go/profile"
	"frontend_go/records"
	"frontend_go/store"
	"time"
)

type Service struct {
	DB       *store.DB
	Notifier *notify.Notifier
}

func New(db *store.DB, n *notify.Notifier) *Service { return &Service{DB: db, Notifier: n} }
func (s *Service) RegisterProfile(p profile.Profile) error {
	p.Normalize()
	if err := p.Validate(); err != nil {
		return err
	}
	return s.DB.SaveProfile(p)
}
func (s *Service) SubmitRecord(r records.Record) error {
	if err := r.Validate(); err != nil {
		return err
	}
	r.Status = "submitted"
	return s.DB.SaveRecord(r)
}
func (s *Service) ReviewRecord(id, actor string) error {
	r, e := s.DB.LoadRecord(id)
	if e != nil {
		return e
	}
	if r.Status != "submitted" {
		return fmt.Errorf("cannot review %s", r.Status)
	}
	r.Status = "reviewed"
	if e = s.DB.SaveRecord(r); e != nil {
		return e
	}
	return s.DB.SaveAudit(archive.NewAudit(fmt.Sprintf("audit-%d", time.Now().UnixNano()), "Record", id, "review", actor, "approved"))
}
func (s *Service) PublishRecord(ctx context.Context, id, actor string) error {
	r, e := s.DB.LoadRecord(id)
	if e != nil {
		return e
	}
	if r.Status != "reviewed" {
		return fmt.Errorf("cannot publish %s", r.Status)
	}
	if e = r.Publish(time.Now()); e != nil {
		return e
	}
	if e = s.DB.SaveRecord(r); e != nil {
		return e
	}
	if e = s.DB.SaveEvent(archive.NewEvent(fmt.Sprintf("event-%d", time.Now().UnixNano()), id, "published", actor)); e != nil {
		return e
	}
	p, _ := s.DB.LoadProfile(r.ProfileID)
	msg := fmt.Sprintf("published %s", r.Title)
	if p.NotificationAddress != "" {
		return s.Notifier.Send(ctx, msg)
	}
	return s.Notifier.Send(ctx, msg)
}
func (s *Service) ArchiveRecord(id, actor string) error {
	r, e := s.DB.LoadRecord(id)
	if e != nil {
		return e
	}
	if e = archive.CanArchive(r); e != nil {
		return e
	}
	if e = r.Archive(); e != nil {
		return e
	}
	if e = s.DB.SaveRecord(r); e != nil {
		return e
	}
	return s.DB.SaveEvent(archive.NewEvent(fmt.Sprintf("event-%d", time.Now().UnixNano()), id, "archived", actor))
}
func (s *Service) Search(f records.Filter) ([]records.Record, error) {
	rs, e := s.DB.AllRecords()
	if e != nil {
		return nil, e
	}
	out := []records.Record{}
	for _, r := range rs {
		if records.Match(r, f) {
			out = append(out, r)
		}
	}
	return records.Paginate(records.SortNewest(out), 0, f.Limit), nil
}
