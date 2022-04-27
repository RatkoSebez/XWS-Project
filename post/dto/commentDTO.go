package dto

type CommentDTO struct {
	CommentContent string `json:"commentContent"`
	PostID         uint   `json:"postId"`
}
