package dto

type UserDTO struct {
	//	ID         int      `json:"id"`
	Name       string   `bson:"name"`
	Surname    string   `bson:"surname"`
	Email      string   `bson:"email"`
	Numtel     string   `bson:"numtel"`
	Sex        string   `bson:"sex"`
	BDateDay   int      `bson:"bdateday"`
	BDateMonth int      `bson:"bdatemonth"`
	BDateYear  int      `bson:"bdateyear"`
	Username   string   `bson:"username"`
	Password   string   `bson:"password"`
	Bio        string   `bson:"bio"`
	Experience []string `bson:"experience"`
	Education  []string `bson:"education"`
	Interests  []string `bson:"interests"`
	Skills     []string `bson:"skills"`
	IsPrivate  bool     `bson:"isPrivate"`
}
