package dto

import (
	"XWS-Project/post/model"
)

type ReactionDTO struct {
	Reaction model.ReactionType `json:"reactionType"`
	UserID   uint               `json:"userId"`
	PostID   uint               `json:"postId"`
}
