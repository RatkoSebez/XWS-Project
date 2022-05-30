package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOffer struct {
	JobOfferId     primitive.ObjectID `bson:"_id" json:"jobOfferId"`
	CompanyId      string             `bson:"companyId" json:"companyId"`
	Position       string             `bson:"position" json:"position"`
	JobDescription string             `bson:"jobDescription" json:"jobDescription"`
	Preconditions  []string           `bson:"preconditions" json:"preconditions"`
}
