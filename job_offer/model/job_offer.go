package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOffer struct {
	JobOfferId     primitive.ObjectID `bson:"_id" json:"jobOfferId"`
	CompanyId      primitive.ObjectID `bson:"companyId" json:"companyId"`
	position       string             `bson:"position" json:"position"`
	jobDescription string             `bson:"jobDescription" json:"jobDescription"`
	preconditions  []string           `bson:"preconditions" json:"preconditions"`
}
