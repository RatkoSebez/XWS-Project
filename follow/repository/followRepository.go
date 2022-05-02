package repository

import (
	"XWS-Project/follow/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FollowRepository struct {
	Client *mongo.Client
}

func (repo *FollowRepository) GetFollow(ctx context.Context, email string) (*model.Follow, error) {
	follow := &model.Follow{}
	collection := repo.Client.Database("dislinkt").Collection("userFollows")
	// fmt.Print(collection.Name())

	filter := bson.D{{"email", email}}
	collection.FindOne(ctx, filter).Decode(&follow)
	return follow, nil
}

func (repo *FollowRepository) CreateFollowRequest(ctx context.Context, followRequest model.FollowRequest) error {
	collection := repo.Client.Database("dislinkt").Collection("followRequest")
	collection.InsertOne(ctx, followRequest)
	return nil
}

func (repo *FollowRepository) GetFollowRequest(ctx context.Context, email string) ([]*model.FollowRequest, error) {
	var result []*model.FollowRequest
	collection := repo.Client.Database("dislinkt").Collection("followRequest")
	//fmt.Println("Repo: " + email)
	filter := bson.D{{"receiveremail", email}}
	cur, _ := collection.Find(ctx, filter)

	// Iterate through the cursor
	for cur.Next(ctx) {
		var elem model.FollowRequest
		cur.Decode(&elem)
		result = append(result, &elem)
	}

	cur.Close(ctx)

	//fmt.Println(len(result))

	return result, nil
}
