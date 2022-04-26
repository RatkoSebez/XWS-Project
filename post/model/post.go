package model

import (
	"time"
)

type Post struct {
	PostID       uint       `json:"postId"`
	UserID       uint       `json:"userId"`
	CreationTime time.Time  `json:"creationTime"`
	MediaAssets  []Media    `json:"mediaAssets"`
	PostText     string     `json:"postText"`
	Reactions    []Reaction `json:"reactions"`
	Comments     []Comment  `json:"comments"`
}
