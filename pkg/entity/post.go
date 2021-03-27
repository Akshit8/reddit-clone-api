// Package entity defines entities and associated functionality for all api resources
package entity

import "time"

// Post defines all fields on entity post.
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


