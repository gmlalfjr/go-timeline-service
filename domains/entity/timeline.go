package entity

import "time"

type Timeline struct {
	ID        int       `json:"id"`
	PostText  string    `json:"post_text"`
	IsPrivate bool      `json:"is_private"`
	CreatedAt time.Time `json:"created_at"`
}
