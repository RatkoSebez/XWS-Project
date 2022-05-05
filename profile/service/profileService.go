package service

import (
	"XWS-Project/profile/dto"
	"XWS-Project/profile/model"
	"XWS-Project/profile/repository"

	"context"
)

type ProfileService struct {
	ProfileRepository *repository.ProfileRepository
}

func (service *ProfileService) UpdateProfile(ctx context.Context, dto dto.ProfileDTO, tempUsername string) error {
	user := model.ProfileInfo{Name: dto.Name, Surname: dto.Surname, Email: dto.Email, Numtel: dto.Numtel, Sex: dto.Sex, BDateDay: dto.BDateDay, BDateMonth: dto.BDateMonth, BDateYear: dto.BDateYear, Username: tempUsername, Password: dto.Password, Bio: dto.Bio, Experience: dto.Experience, Education: dto.Education, Interests: dto.Interests, Skills: dto.Skills, IsPrivate: dto.IsPrivate}
	err := service.ProfileRepository.UpdateUser(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}
func (service *ProfileService) FindUserByUsername(ctx context.Context, username string, tempUsername string) *model.ProfileInfo {
	user := service.ProfileRepository.FindUserByUsername(ctx, username, tempUsername)

	return user
}
func (service *ProfileService) FindUserByMail(ctx context.Context, email string) *model.ProfileInfo {
	user := service.ProfileRepository.FindUserByMail(ctx, email)
	if user == nil {
		return nil
	}
	return user
}
