package entity

import "time"

type Comment struct {
	ID        int       `json:"id"`
	PostText  string    `json:"comment_text"`
	CreatedAt time.Time `json:"created_at"`
}
