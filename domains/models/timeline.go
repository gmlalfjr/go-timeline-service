package models

type TimelineRequest struct {
	PostText  string `json:"post_text"`
	IsPrivate bool   `json:"is_private"`
}
