package service

import (
	"XWS-Project/registration/model"
	"XWS-Project/registration/repository"
	"context"

	"XWS-Project/registration/dto"
)

type RegistrationService struct {
	RegistrationRepository *repository.RegistrationRepository
}

func (service *RegistrationService) RegisterUser(ctx context.Context, dto dto.UserDTO) error {
	user := model.User{ID: dto.ID, Name: dto.Name, Surname: dto.Surname, Email: dto.Email, Numtel: dto.Numtel, Sex: dto.Sex, BDateDay: dto.BDateDay, BDateMonth: dto.BDateMonth, BDateYear: dto.BDateYear, Username: dto.Username, Password: dto.Password, Bio: dto.Bio, Experience: dto.Experience, Interests: dto.Interests, IsPrivate: dto.IsPrivate}
	err := service.RegistrationRepository.RegisterUser(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

func (service *RegistrationService) FindUserByUsername(ctx context.Context, username string) *model.User {
	user := service.RegistrationRepository.FindUserByUsername(ctx, username)

	return user
}
