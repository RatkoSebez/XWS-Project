package repository

import (
	"XWS-Project/post/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	collection := repository.Client.Database("dislinkt").Collection("posts")
	findOptions := options.Find()
	findOptions.SetLimit(30)
	filter := bson.D{{"userId", userID}}
	var results []*model.Post

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		panic(err)
	}

	for cur.Next(ctx) {
		var elem model.Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	cur.Close(ctx)
	return results, err
}

func (repository *PostRepository) SaveMedia(ctx context.Context, data model.Media) error {
	collection := repository.Client.Database("dislinkt").Collection("media")
	result, err := collection.InsertOne(ctx, data)
	fmt.Println("Inserted media", result.InsertedID)
	return err
}

func (repository *PostRepository) LoadPostByID(ctx context.Context, id uint) (*model.Post, error) {
	filter := bson.D{{"postId", id}}
	var result model.Post
	collection := repository.Client.Database("dislinkt").Collection("posts")
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result, nil
}

func (repository *PostRepository) UpdatePost(ctx context.Context, post model.Post) (*model.Post, error) {
	filter := bson.D{{"postId", post.PostID}}
	update := bson.D{{"$set", bson.D{{"comments", post.Comments}}}}
	collection := repository.Client.Database("dislinkt").Collection("posts")
	_, err := collection.UpdateOne(ctx, filter, update)
	return nil, err
}

func (repository *PostRepository) AddCommentToPost(ctx context.Context, postId uint, comment model.Comment) (*model.Post, error) {
	collection := repository.Client.Database("dislinkt").Collection("comments")
	rslt, err := collection.InsertOne(ctx, comment)
	fmt.Println("Inserted comment", rslt.InsertedID)
	if err != nil {
		panic(err)
	}
	result, _ := repository.LoadPostByID(ctx, postId)
	result.Comments = append(result.Comments, comment)
	updateRes, err := repository.UpdatePost(ctx, *result)
	return updateRes, err
}

func (repository *PostRepository) ReactOnPost(ctx context.Context, post model.Post, reaction model.Reaction) (*model.Post, error) {
	collection := repository.Client.Database("dislinkt").Collection("reactions")
	_, err := collection.InsertOne(ctx, reaction)
	if err != nil {
		panic(err)
	}
	updateRes, err := repository.UpdatePost(ctx, post)
	return updateRes, err
}
