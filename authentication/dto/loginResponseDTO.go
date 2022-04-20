package dto

type LoginResponseDTO struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	ProfileID uint   `json:"profileID"`
	Username  string `json:"username"`
}
