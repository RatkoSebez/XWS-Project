package repository

import (
	"XWS-Project/post/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostRepository struct {
	Client *mongo.Client
}

func (repository *PostRepository) MakePost(ctx context.Context, data model.Post) (*model.Post, error) {
	collection := repository.Client.Database("dislinkt").Collection("posts")
	result, err := collection.InsertOne(ctx, data)
	fmt.Println("Inserted post", result.InsertedID)
	return &data, err
}

func (repository *PostRepository) LoadPosts(ctx context.Context) ([]*model.Post, error) {
	return nil, nil
}

func (repository *PostRepository) LoadPostsByUser(ctx context.Context, userEmail string) ([]*model.Post, error) {
	collection := repository.Client.Database("dislinkt").Collection("posts")
	findOptions := options.Find()
	findOptions.SetLimit(30)
	filter := bson.D{{"useremail", userEmail}}
	var results []*model.Post

	cur, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		panic(err)
	}

	for cur.Next(ctx) {
		var elem model.Post
		err := cur.Decode(&elem)
		if err != nil {
			panic(err)
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

func (repository *PostRepository) LoadPostByID(ctx context.Context, id primitive.ObjectID) (*model.Post, error) {
	filter := bson.D{{"_id", id}}
	var result model.Post
	collection := repository.Client.Database("dislinkt").Collection("posts")
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return &result, nil
}

func (repository *PostRepository) UpdatePost(ctx context.Context, post model.Post) (*model.Post, error) {
	filter := bson.D{{"_id", post.PostID}}
	update := bson.D{
		{"$set", bson.D{
			{"comments", post.Comments},
			{"reactions", post.Reactions},
			{"postText", post.PostText},
		}}}
	collection := repository.Client.Database("dislinkt").Collection("posts")
	result, err := collection.UpdateOne(ctx, filter, update)
	if result.MatchedCount == 0 {
		fmt.Println("No documents found to update")
	}
	return nil, err
}

func (repository *PostRepository) AddCommentToPost(ctx context.Context, postId primitive.ObjectID, comment model.Comment) (*model.Post, error) {
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
	ress, err := collection.InsertOne(ctx, reaction)
	fmt.Println(ress)
	if err != nil {
		panic(err)
	}
	_, err = repository.UpdatePost(ctx, post)
	updateRes, _ := repository.LoadPostByID(ctx, post.PostID)
	return updateRes, err
}
