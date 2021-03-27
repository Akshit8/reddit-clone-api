// Package entity defines all api entities.
package entity

import "time"

// Post defines entity POST.
type Post struct {
	ID        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
