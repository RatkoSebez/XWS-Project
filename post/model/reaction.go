package model

import (
	"time"
)

type Reaction struct {
	UserID       uint      `json:"userId"`
	CreationTime time.Time `json:"creationTime"`
	ReactionType string    `json:"reactionType"`
}
