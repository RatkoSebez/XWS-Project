package repository

import (
	"XWS-Project/follow/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowRepository struct {
	Client *mongo.Client
}

func (repo *FollowRepository) Follow(ctx context.Context, email string) (*model.Follow, error) {
	follow := &model.Follow{}
	collection := repo.Client.Database("dislinkt").Collection("userFollows")
	// fmt.Print(collection.Name())

	filter := bson.D{{"email", email}}
	collection.FindOne(ctx, filter).Decode(&follow)
	return follow, nil
}

func (repo *FollowRepository) CreateFollowRequest(ctx context.Context, followRequest model.FollowRequest) error {
	fmt.Println(followRequest)
	collection := repo.Client.Database("dislinkt").Collection("followRequest")
	collection.InsertOne(ctx, followRequest)
	return nil
}
