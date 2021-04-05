// Package entity defines entities and associated functionality for all api resources
package entity

import (
	"time"
)

// Post defines all fields on entity post.
type Post struct {
	ID        int       `json:"id"`
	Owner     int       `json:"owner"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UpVotes   int       `json:"upvotes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
