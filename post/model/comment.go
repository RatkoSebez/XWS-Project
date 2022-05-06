package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	CommentID    primitive.ObjectID `bson:"_id" json:"commentId"`
	PostID       primitive.ObjectID `json:"postId"`
	UserEmail    string             `json:"userEmail"`
	CreationTime time.Time          `json:"creationTime"`
	CommentData  string             `json:"comment"`
}
