package dto

type NewPostDTO struct {
	Links    []string `json:"links"`
	Photos   []string `json:"photos"`
	PostText string   `json:"postText"`
}
