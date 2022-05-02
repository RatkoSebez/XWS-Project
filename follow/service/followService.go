package service

import (
	"XWS-Project/follow/model"
	"XWS-Project/follow/repository"
	"context"
	"fmt"
)

type FollowService struct {
	FollowRepo *repository.FollowRepository
}

func (service *FollowService) GetFollow(ctx context.Context, email string) (*model.Follow, error) {
	userFollows, err := service.FollowRepo.Follow(ctx, email)
	if err != nil {
		fmt.Print(err)
	}
	return userFollows, nil
}

func (service *FollowService) CreateFollowRequest(ctx context.Context, followRequest model.FollowRequest) error {
	err := service.FollowRepo.CreateFollowRequest(ctx, followRequest)
	return err
}
