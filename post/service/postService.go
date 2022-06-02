package service

import (
	"XWS-Project/post/dto"
	"XWS-Project/post/model"
	"XWS-Project/post/repository"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostService struct {
	PostRepo *repository.PostRepository
}

func (service *PostService) MakePost(ctx context.Context, data dto.NewPostDTO, userEmail string) (*model.Post, error) {
	var post model.Post
	post.PostID = primitive.NewObjectID()
	post.UserEmail = userEmail
	post.Reactions = []model.Reaction{}
	post.PostText = data.PostText
	post.Comments = []model.Comment{}
	post.MediaAssets = []model.Media{}

	for _, el := range data.Links {
		var asset model.Media
		asset.MediaID = primitive.NewObjectID()
		asset.Site = el
		asset.Filepath = ""
		err := service.PostRepo.SaveMedia(ctx, asset)
		if err != nil {
			panic(err)
		}
		post.MediaAssets = append(post.MediaAssets, asset)
	}

	idString := userEmail

	for _, el := range data.Photos {
		var asset model.Media
		asset.MediaID = primitive.NewObjectID()
		asset.Site = ""
		asset.Filepath = "../photos/" + idString + el
		err := service.PostRepo.SaveMedia(ctx, asset)
		if err != nil {
			panic(err)
		}
		post.MediaAssets = append(post.MediaAssets, asset)
	}
	post.CreationTime = time.Now()
	resp, err := service.PostRepo.MakePost(ctx, post)

	return resp, err
}

func (service *PostService) SavePhotos(request dto.PhotoDTO, userEmail string) {
	var fileData = request.FileData
	idx := strings.Index(fileData, ";base64,")
	ImageType := fileData[11:idx]
	log.Println(ImageType)
	unbased, err := base64.RawStdEncoding.DecodeString(fileData[idx+8 : len(fileData)-1])
	if err != nil {
		panic(err)
	}
	byteReader := bytes.NewReader(unbased)
	idString := userEmail
	switch ImageType {
	case "png":
		im, err := png.Decode(byteReader)
		if err != nil {
			panic(err)
		}

		f, err := os.OpenFile("../photos/"+idString+request.FileName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic(err)
		}

		png.Encode(f, im)
	case "jpeg":
		im, err := jpeg.Decode(byteReader)
		if err != nil {
			panic(err)
		}

		f, err := os.OpenFile("../photos/"+idString+request.FileName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic(err)
		}

		jpeg.Encode(f, im, nil)
	case "gif":
		im, err := gif.Decode(byteReader)
		if err != nil {
			panic(err)
		}

		f, err := os.OpenFile("../photos/"+idString+request.FileName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic(err)
		}

		gif.Encode(f, im, nil)
	}
}
func (service *PostService) PostComment(ctx context.Context, data dto.CommentDTO, userEmail string) (*model.Post, error) {
	var comm model.Comment
	comm.CommentID = primitive.NewObjectID()
	comm.CreationTime = time.Now()
	comm.PostID = data.PostID
	comm.CommentData = data.CommentContent
	comm.UserEmail = userEmail
	post, err := service.PostRepo.AddCommentToPost(ctx, data.PostID, comm)
	return post, err
}

func (service *PostService) MakeReaction(ctx context.Context, data dto.ReactionDTO) (*model.Post, error) {
	post, err := service.PostRepo.LoadPostByID(ctx, data.PostID)
	if err != nil {
		panic(err)
	}
	var reaction model.Reaction
	reaction.ReactionID = primitive.NewObjectID()
	reaction.PostID = data.PostID
	reaction.UserEmail = data.UserEmail
	reaction.CreationTime = time.Now()
	reaction.TypeOfReaction = data.React
	post.Reactions = append(post.Reactions, reaction)
	result, err := service.PostRepo.ReactOnPost(ctx, *post, reaction)
	return result, err
}

func (service *PostService) LoadFollowingPosts(ctx context.Context, followingList []string) ([]*model.Post, error) {
	var returnList []*model.Post
	for _, el := range followingList {
		list, err := service.PostRepo.LoadPostsByUser(ctx, el)
		if err != nil {
			panic(err)
		}
		for _, element := range list {
			returnList = append(returnList, element)
		}
	}

	return returnList, nil
}

func (service *PostService) GetUserPosts(ctx context.Context, userEmail string) ([]*model.Post, error) {

	list, err := service.PostRepo.LoadPostsByUser(ctx, userEmail)
	if len(list) == 0 {
		fmt.Println("Nema postova")
	}
	if len(list) != 0 {
		fmt.Println("Ima postova")
		fmt.Println(list[0].PostText)
	}
	return list, err
}
