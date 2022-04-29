package dto

type UserDTO struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	Surname    string   `json:"surname"`
	Email      string   `json:"email"`
	Numtel     string   `json:"numtel"`
	Sex        string   `json:"sex"`
	BDateDay   int      `json:"bdateday"`
	BDateMonth int      `json:"bdatemonth"`
	BDateYear  int      `json:"bdateyear"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	Bio        string   `json:"bio"`
	Experience []string `json:"experience"`
	Interests  []string `json:"interests"`
	IsPrivate  bool     `json:"isPrivate"`
}
