package model

import (
	"time"
)

type Reaction struct {
	ReactionID     uint         `json:"reactionId"`
	PostID         uint         `json:"postId"`
	UserID         uint         `json:"userId"`
	CreationTime   time.Time    `json:"creationTime"`
	TypeOfReaction ReactionType `json:"reactionType"`
}
