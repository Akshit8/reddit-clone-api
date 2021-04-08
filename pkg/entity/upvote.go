package entity

import "time"

// Upvote defines all fields on entity upvote.
type Upvote struct {
	UserID    int       `json:"userId"`
	PostID    int       `json:"postId"`
	Value     int       `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
