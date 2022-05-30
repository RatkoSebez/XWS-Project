package model

type Job struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Requirements string `json:"requirements"`
	Benefits     string `json:"benefits"`
}
