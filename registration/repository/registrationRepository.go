package repository

import (
	"XWS-Project/registration/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type RegistrationRepository struct {
	Client *mongo.Client
}

func (repository *RegistrationRepository) FindUserByUsername(ctx context.Context, username string) *model.User {
	user := &model.User{}
	collection := repository.Client.Database("dislinkt").Collection("users")
	filter2 := bson.D{{"username", username}}

	err2 := collection.FindOne(ctx, filter2).Decode(&user)
	if err2 != nil {

		return nil
	}

	return user

}

func (repository *RegistrationRepository) RegisterUser(ctx context.Context, user *model.User) error {
	collection := repository.Client.Database("dislinkt").Collection("users")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User registrated: ", insertResult.InsertedID)

	return nil
}
