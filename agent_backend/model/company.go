package model

type Company struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	Description string `json:"description"`
	OwnerEmail  string `json:"ownerEmail" bson:"ownerEmail"`
	IsApproved  bool   `json:"isApproved"`
}
