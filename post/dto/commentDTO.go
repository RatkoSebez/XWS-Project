package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentDTO struct {
	CommentContent string             `json:"commentContent"`
	PostID         primitive.ObjectID `json:"postId"`
}
