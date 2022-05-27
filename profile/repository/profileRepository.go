package repository

import (
	"XWS-Project/profile/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileRepository struct {
	Client *mongo.Client
}

func (repository *ProfileRepository) UpdateUser(ctx context.Context, user *model.ProfileInfo) error {
	collection := repository.Client.Database("dislinkt").Collection("users")
	filter := bson.D{{"email", user.Email}}
	update := bson.D{{"$set", bson.D{
		{"name", user.Name}, {"surname", user.Surname}, {"email", user.Email}, {"numtel", user.Numtel}, {"sex", user.Sex}, {"bdatemonth", user.BDateMonth}, {"bdateday", user.BDateDay}, {"bdateyear", user.BDateYear}, {"username", user.Username}, {"password", user.Password}, {"bio", user.Bio}, {"experience", user.Experience}, {"education", user.Education}, {"interests", user.Interests}, {"skills", user.Skills}, {"isPrivate", user.IsPrivate},
	}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

func (repository *ProfileRepository) FindUserByUsername(ctx context.Context, username string, tempUsername string) *model.ProfileInfo {
	user := &model.ProfileInfo{}
	collection := repository.Client.Database("dislinkt").Collection("users")
	filter2 := bson.D{{"username", username}}
	err2 := collection.FindOne(ctx, filter2).Decode(&user)
	if err2 != nil {
		return nil
	}
	return user
}
func (repository *ProfileRepository) FindUserByMail(ctx context.Context, mail string) *model.ProfileInfo {
	user := &model.ProfileInfo{}
	collection := repository.Client.Database("dislinkt").Collection("users")
	filter2 := bson.D{{"email", mail}}
	err2 := collection.FindOne(ctx, filter2).Decode(&user)
	if err2 != nil {
		return nil
	}
	return user
}

func (repository *ProfileRepository) GetAllUsers(ctx context.Context) []*model.ProfileInfo {
	collection := repository.Client.Database("dislinkt").Collection("users")
	cur, err := collection.Find(ctx, nil)
	var results []*model.ProfileInfo

	if err != nil {
		panic(err)
	}

	for cur.Next(ctx) {
		var elem model.ProfileInfo
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
		}

		results = append(results, &elem)
	}

	cur.Close(ctx)
	return results
}
