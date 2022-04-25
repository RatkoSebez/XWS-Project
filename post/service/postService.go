package service

import (
	"XWS-Project/post/dto"
	"XWS-Project/post/model"
	"XWS-Project/post/repository"
	"context"
	"log"
	"time"
)

type PostService struct {
	PostRepo *repository.PostRepository
}

func (service *PostService) MakePost(ctx context.Context, data dto.NewPostDTO, userID uint) (*model.Post, error) {
	var post model.Post
	post.UserID = userID
	post.Reactions = []model.Reaction{}
	post.PostText = data.PostText

	for _, el := range data.Links {
		var asset model.Media
		asset.Site = el
		asset.Filepath = ""
		err := service.PostRepo.SaveMedia(ctx, asset)
		if err != nil {
			log.Fatal(err)
		}
		post.MediaAssets = append(post.MediaAssets, asset)
	}

	for _, el := range data.Photos {
		var asset model.Media
		asset.Site = ""
		asset.Filepath = el
		err := service.PostRepo.SaveMedia(ctx, asset)
		if err != nil {
			log.Fatal(err)
		}
		post.MediaAssets = append(post.MediaAssets, asset)
	}
	post.CreationTime = time.Now()
	resp, err := service.PostRepo.MakePost(ctx, post)

	return resp, err
}

func (service *PostService) SavePhotos() {

}
