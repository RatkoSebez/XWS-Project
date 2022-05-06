package repository

import (
	"XWS-Project/authentication/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthenticationRepository struct {
	Client *mongo.Client
}

func (repo *AuthenticationRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	collection := repo.Client.Database("dislinkt").Collection("users")
	fmt.Print(collection.Name())

	filter := bson.D{{"email", email}}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		panic(err)
	}

	return user, nil
}
