package dto

type RegistrationResponseDTO struct { //proveri
	Mail     string `json:"mail"`
	Token    string `json:"token"`
	ID       uint   `json:"id"`
	Username string `json:"username"`
}
