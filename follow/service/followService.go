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
	userFollows, err := service.FollowRepo.GetFollow(ctx, email)
	if err != nil {
		fmt.Print(err)
	}
	return userFollows, nil
}

func (service *FollowService) CreateFollowRequest(ctx context.Context, followRequest model.FollowRequest) error {
	err := service.FollowRepo.CreateFollowRequest(ctx, followRequest)
	return err
}

func (service *FollowService) GetFollowRequest(ctx context.Context, email string) ([]*model.FollowRequest, error) {
	followRequests, err := service.FollowRepo.GetFollowRequest(ctx, email)
	if err != nil {
		fmt.Print(err)
	}
	return followRequests, nil
}
