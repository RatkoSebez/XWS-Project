package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Reaction struct {
	ReactionID     primitive.ObjectID `bson:"_id" json:"reactionId"`
	PostID         primitive.ObjectID `json:"postId"`
	UserEmail      string             `json:"userEmail"`
	CreationTime   time.Time          `json:"creationTime"`
	TypeOfReaction ReactionType       `json:"reactionType"`
}
