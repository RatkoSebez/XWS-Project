package service

import (
	"XWS-Project/post/dto"
	"XWS-Project/post/model"
	"XWS-Project/post/repository"
	"bytes"
	"context"
	"encoding/base64"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
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
	post.Comments = []model.Comment{}

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

	idString := strconv.FormatUint(uint64(userID), 10)

	for _, el := range data.Photos {
		var asset model.Media
		asset.Site = ""
		asset.Filepath = "../photos/" + idString + el
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

func (service *PostService) SavePhotos(request dto.PhotoDTO, userID uint) {
	var fileData = request.FileData
	idx := strings.Index(fileData, ";base64,")
	ImageType := fileData[11:idx]
	log.Println(ImageType)
	unbased, err := base64.RawStdEncoding.DecodeString(fileData[idx+8 : len(fileData)-1])
	if err != nil {
		log.Fatal(err)
	}
	byteReader := bytes.NewReader(unbased)
	idString := strconv.FormatUint(uint64(userID), 10)
	switch ImageType {
	case "png":
		im, err := png.Decode(byteReader)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile("../photos/"+idString+request.FileName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}

		png.Encode(f, im)
	case "jpeg":
		im, err := jpeg.Decode(byteReader)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile("../photos/"+idString+request.FileName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}

		jpeg.Encode(f, im, nil)
	case "gif":
		im, err := gif.Decode(byteReader)
		if err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile("../photos/"+idString+request.FileName, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}

		gif.Encode(f, im, nil)
	}
}
func (service *PostService) PostComment() (*model.Post, error) {
	return nil, nil
}
func (service *PostService) Like() (*model.Post, error) {
	return nil, nil
}
func (service *PostService) Dislike() (*model.Post, error) {
	return nil, nil
}
func (service *PostService) Unlike() (*model.Post, error) {
	return nil, nil
}
func (service *PostService) Undislike() (*model.Post, error) {
	return nil, nil
}
