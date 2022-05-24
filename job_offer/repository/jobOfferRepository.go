package repository

import (
	"XWS-Project/job_offer/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobRepository struct {
	Client *mongo.Client
}

func (repository *JobRepository) AddNew(ctx context.Context, data model.JobOffer) (*model.JobOffer, error) {
	collection := repository.Client.Database("jobOffers").Collection("offers")
	result, err := collection.InsertOne(ctx, data)
	fmt.Println("Inserted post", result.InsertedID)
	return &data, err
}

func (repository *JobRepository) FindByCompany(ctx context.Context, data primitive.ObjectID) ([]*model.JobOffer, error) {
	collection := repository.Client.Database("jobOffers").Collection("offers")
	findOptions := options.Find()
	findOptions.SetLimit(30)
	filter := bson.D{{"companyId", data}}
	var results []*model.JobOffer

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		panic(err)
	}

	for cur.Next(ctx) {
		var elem model.JobOffer
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		results = append(results, &elem)
	}

	cur.Close(ctx)
	return results, err
}

func (repository *JobRepository) FindByPosition(ctx context.Context, data string) ([]*model.JobOffer, error) {
	collection := repository.Client.Database("jobOffers").Collection("offers")
	findOptions := options.Find()
	findOptions.SetLimit(30)
	filter := bson.D{{"position", data}}
	var results []*model.JobOffer

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		panic(err)
	}

	for cur.Next(ctx) {
		var elem model.JobOffer
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		results = append(results, &elem)
	}

	cur.Close(ctx)
	return results, err
}
