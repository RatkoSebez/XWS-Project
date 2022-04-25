package dto

type NewPostDTO struct {
	UserID   uint     `json:"userId"`
	Links    []string `json:"links"`
	Photos   []string `json:"photos"`
	PostText string   `json:"postText"`
}
