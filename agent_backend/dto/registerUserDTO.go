package dto

type RegisterUserDTO struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Numtel   string `json:"numtel"`
	Username string `json:"username"`
	Password string `json:"password"`
}
