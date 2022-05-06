package dto

type RegisterDTO struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}
