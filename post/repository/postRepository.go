package repository

import (
	"XWS-Project/post/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct {
	Client *mongo.Client
}

func (repository *PostRepository) MakePost(ctx context.Context, data model.Post) (*model.Post, error) {
	collection := repository.Client.Database("dislinkt").Collection("posts")
	result, err := collection.InsertOne(ctx, data)
	fmt.Println("Inserted media", result.InsertedID)
	return &data, err
}

func (repository *PostRepository) LoadPosts(ctx context.Context) ([]*model.Post, error) {
	return nil, nil
}

func (repository *PostRepository) LoadPostsByUser(ctx context.Context, userID uint) ([]*model.Post, error) {
	return nil, nil
}

func (repository *PostRepository) SaveMedia(ctx context.Context, data model.Media) error {
	collection := repository.Client.Database("dislinkt").Collection("media")
	result, err := collection.InsertOne(ctx, data)
	fmt.Println("Inserted media", result.InsertedID)
	return err
}
