package model

type Media struct {
	MediaID  uint   `json:"mediaID"`
	Site     string `json:"site"`
	Filepath string `json:"filepath"`
}
