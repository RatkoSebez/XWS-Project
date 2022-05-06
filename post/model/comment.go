package model

import (
	"time"
)

type Comment struct {
	CommentID    uint      `json:"commentId"`
	PostID       uint      `json:"postId"`
	UserID       uint      `json:"userId"`
	CreationTime time.Time `json:"creationTime"`
	CommentData  string    `json:"comment"`
}
