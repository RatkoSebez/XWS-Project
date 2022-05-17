package model

type ReactionType string

const (
	Like    ReactionType = "LIKE"
	Dislike ReactionType = "DISLIKE"
)

func (rt ReactionType) String() string {
	if rt == Like {
		return "LIKE"
	}
	if rt == Dislike {
		return "DISLIKE"
	}
	return ""
}
