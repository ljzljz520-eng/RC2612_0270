package api

import (
	"encoding/json"
	"frontend_go/notify"
	"frontend_go/profile"
	"frontend_go/records"
	"frontend_go/store"
	"frontend_go/workflow"
	"net/http"
	"strings"
)

type Service struct{ Core *workflow.Service }

func NewService(db *store.DB, n *notify.Notifier) *Service {
	return &Service{Core: workflow.New(db, n)}
}
func Routes(s *Service) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) })
	mux.HandleFunc("/profiles", s.profiles)
	mux.HandleFunc("/records", s.records)
	return mux
}
func decode(w http.ResponseWriter, r *http.Request, v any) bool {
	if e := json.NewDecoder(r.Body).Decode(v); e != nil {
		http.Error(w, e.Error(), 400)
		return false
	}
	return true
}
func reply(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
func (s *Service) profiles(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var p profile.Profile
		if !decode(w, r, &p) {
			return
		}
		if e := s.Core.RegisterProfile(p); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		reply(w, p)
		return
	}
	http.Error(w, "method not allowed", 405)
}
func (s *Service) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var rec records.Record
		if !decode(w, r, &rec) {
			return
		}
		if e := s.Core.SubmitRecord(rec); e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		reply(w, rec)
		return
	}
	if r.Method == http.MethodGet {
		q := r.URL.Query().Get("q")
		rs, e := s.Core.Search(records.Filter{Query: strings.ToLower(q), Limit: 50})
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
		reply(w, rs)
		return
	}
	http.Error(w, "method not allowed", 405)
}
