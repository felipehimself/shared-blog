package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	Id              uint64 `json:"id,omitempty"`
	Username        string `json:"username,omitempty"`
	Topic           string `json:"topic,omitempty"`
	TopicId         uint64 `json:"topicId,omitempty"`
	Title           string `json:"title,omitempty"`
	Subtitle        string `json:"subtitle,omitempty"`
	Author          string `json:"author,omitempty"`
	AuthorId        uint64 `json:"authorId,omitempty"`
	Content         string `json:"content,omitempty"`
	Votes           uint64 `json:"votes"`
	Voted           bool   `json:"voted"`
	Comments        uint64 `json:"comments"`
	MinutesRead     uint64 `json:"minutesRead,omitempty"`
	CommentsContent []Comment `json:"commentsContent,omitempty"`
	CreatedAt       time.Time `json:"createdAt,omitempty"`
}

func (p *Post) ValidatePostFields() error {
	p.cleanWhiteSpaces()

	if p.Title == "" && p.Subtitle == "" && p.Content == "" && p.TopicId == 0 {
		return errors.New("all fields are required")
	}

	if p.TopicId == 0 {
		return errors.New("topic id is required")
	}

	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Subtitle == "" {
		return errors.New("subtitle is required")
	}

	if p.Content == "" {
		return errors.New("content is required")
	}

	return nil

}

func (p *Post) cleanWhiteSpaces() {

	p.Title = strings.TrimSpace(p.Title)
	p.Subtitle = strings.TrimSpace(p.Subtitle)
	p.Content = strings.TrimSpace(p.Content)

}
