package dto

type ProfileDTO struct {
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
	Education  []string `json:"education"`
	Interests  []string `json:"interests"`
	Skills     []string `json:"skills"`
	IsPrivate  bool     `json:"isPrivate"`
}
