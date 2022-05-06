package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Media struct {
	MediaID  primitive.ObjectID `bson:"_id" json:"mediaId"`
	Site     string             `json:"site"`
	Filepath string             `json:"filepath"`
}
