package records

import (
	"errors"
	"strings"
	"time"
)

type Record struct {
	ID, ProfileID, Kind, Title, Abstract, Status string
	Authors                                      []string
	Keywords                                     []string
	CreatedAt, PublishedAt                       time.Time
	Version                                      int
}

func New(id, profileID, title string) Record {
	return Record{ID: id, ProfileID: profileID, Title: title, Status: "draft", Version: 1, CreatedAt: time.Now().UTC()}
}
func (r Record) Validate() error {
	if r.ID == "" || r.ProfileID == "" || strings.TrimSpace(r.Title) == "" {
		return errors.New("record identity and title required")
	}
	if r.Kind == "" {
		return errors.New("record kind required")
	}
	return nil
}
func (r *Record) SetAbstract(text string) { r.Abstract = strings.TrimSpace(text); r.Version++ }
func (r *Record) AddAuthor(author string) {
	author = strings.TrimSpace(author)
	if author != "" {
		r.Authors = append(r.Authors, author)
		r.Version++
	}
}
func (r *Record) AddKeyword(word string) {
	word = strings.ToLower(strings.TrimSpace(word))
	if word != "" {
		r.Keywords = append(r.Keywords, word)
		r.Version++
	}
}
func (r *Record) Publish(at time.Time) error {
	if r.Status == "archived" {
		return errors.New("archived record cannot publish")
	}
	if len(r.Abstract) < 20 {
		return errors.New("abstract too short")
	}
	r.Status = "published"
	r.PublishedAt = at.UTC()
	r.Version++
	return nil
}
func (r *Record) Archive() error {
	if r.Status != "published" {
		return errors.New("only published records can archive")
	}
	r.Status = "archived"
	r.Version++
	return nil
}
func (r Record) SearchText() string {
	return strings.ToLower(strings.Join([]string{r.Title, r.Abstract, strings.Join(r.Authors, " "), strings.Join(r.Keywords, " ")}, " "))
}
