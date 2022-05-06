package service

import (
	"XWS-Project/authentication/dto"
	"XWS-Project/authentication/model"
	"XWS-Project/authentication/repository"
	"context"
)

type AuthenticationService struct {
	AuthRepo *repository.AuthenticationRepository
}

func (service *AuthenticationService) Login(ctx context.Context, data dto.LoginDTO) (*model.User, error) {
	user, err := service.AuthRepo.GetUserByEmail(ctx, data.Email)
	if err != nil {
		panic(err)
	}
	return user, nil
}
