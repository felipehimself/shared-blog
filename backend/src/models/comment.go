package models

import (
	"errors"
	"strings"
	"time"
)

type Comment struct {
	Id        uint64    `json:"id"`
	UserId    uint64    `json:"userId"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	PostId    uint64    `json:"postId"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}

func (c *Comment) ValidateCommentFields() error {
	c.Comment = strings.TrimSpace(c.Comment)

	if c.Comment == "" && c.PostId == 0 {
		return errors.New("all fields are required")
	}

	if c.Comment == "" {
		return errors.New("comment cannot be empty")

	}

	if c.PostId == 0 {
		return errors.New("post id cannot be empty")
	}

	return nil

}
