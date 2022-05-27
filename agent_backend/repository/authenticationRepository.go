package repository

import (
	"XWS-Project/agent_backend/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AuthenticationRepository struct {
	Client *mongo.Client
}

func (repo *AuthenticationRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	collection := repo.Client.Database("agent").Collection("users")

	filter := bson.D{{"email", email}}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *AuthenticationRepository) RegisterUser(ctx context.Context, user *model.User) error {
	collection := repository.Client.Database("agent").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User registrated: ", insertResult.InsertedID)

	return nil
}
