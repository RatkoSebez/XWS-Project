package model

type Follow struct {
	Email string `json:"email"`
	// email of accounts that are followed
	// think if it is good to store whole user object instead of just email
	Follows   []string `json:"follows"`
	Followers []string `json:"followers"`
}
