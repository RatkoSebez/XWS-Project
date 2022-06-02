package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	PostID       primitive.ObjectID `bson:"_id" json:"postId"`
	UserEmail    string             `bson:"userEmail" json:"userEmail"`
	CreationTime time.Time          `json:"creationTime"`
	MediaAssets  []Media            `json:"mediaAssets"`
	PostText     string             `json:"postText"`
	Reactions    []Reaction         `json:"reactions"`
	Comments     []Comment          `json:"comments"`
}
