package model

type User struct {
	ProfileID uint   `json:"profileId"`
	Email     string `json:"email" gorm:"not null;unique"`
	Username  string `json:"username" gorm:"not null;unique"`
	Password  string `json:"password"`
	Public    bool   `json:"public"`
}
