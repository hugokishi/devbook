package models

import (
	"errors"
	"strings"
	"time"
)

// Publication - Model for publication
type Publication struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	AuthorID  uint64    `json:"author_id,omitempty"`
	Likes     uint64    `json:"likes"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare - Call methods to format and validate recieved publications
func (publication *Publication) Prepare() error {
	if err := publication.validate(); err != nil {
		return err
	}

	publication.format()
	return nil
}

// validate - Validate Publication
func (publication *Publication) validate() error {
	if publication.Title == "" {
		return errors.New("The title is required")
	}
	if publication.Content == "" {
		return errors.New("The Content is required")
	}
	return nil
}

// Format - Format Publication (remove spaces)
func (publication *Publication) format() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Content = strings.TrimSpace(publication.Content)
}
