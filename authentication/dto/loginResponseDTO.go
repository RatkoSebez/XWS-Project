package dto

type LoginResponseDTO struct {
	Email    string `json:"email"`
	Token    string `json:"token"`
	Username string `json:"username"`
}
