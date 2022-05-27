package service

import (
	"XWS-Project/agent_backend/dto"
	"XWS-Project/agent_backend/model"
	"XWS-Project/agent_backend/repository"
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

func (service *AuthenticationService) RegisterUser(ctx context.Context, dto dto.RegisterUserDTO) error {
	user := model.User{Name: dto.Name, Surname: dto.Surname, Email: dto.Email, Username: dto.Username, Password: dto.Password}
	err := service.AuthRepo.RegisterUser(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthenticationService) FindUserByEmail(ctx context.Context, username string) *model.User {
	user, _ := service.AuthRepo.GetUserByEmail(ctx, username)
	return user
}
