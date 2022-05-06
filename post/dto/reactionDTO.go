package dto

import (
	"XWS-Project/post/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReactionDTO struct {
	React     model.ReactionType `json:"reactionType"`
	UserEmail string             `json:"userEmail"`
	PostID    primitive.ObjectID `json:"postId"`
}
