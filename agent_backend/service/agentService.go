package service

import (
	"XWS-Project/agent_backend/dto"
	"XWS-Project/agent_backend/model"
	"XWS-Project/agent_backend/repository"
	"context"
)

type AgentService struct {
	Repository *repository.AgentRepository
}

func (service *AgentService) Login(ctx context.Context, data dto.LoginDTO) (*model.User, error) {
	user, err := service.Repository.GetUserByEmail(ctx, data.Email)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (service *AgentService) RegisterUser(ctx context.Context, dto dto.RegisterUserDTO) error {
	user := model.User{Name: dto.Name, Surname: dto.Surname, Email: dto.Email, Username: dto.Username, Password: dto.Password}
	err := service.Repository.RegisterUser(ctx, &user)
	if err != nil {
		return err
	}
	return nil
}

func (service *AgentService) FindUserByEmail(ctx context.Context, username string) *model.User {
	user, _ := service.Repository.GetUserByEmail(ctx, username)
	return user
}

func (service *AgentService) CreateCompany(ctx context.Context, company model.Company) error {
	err := service.Repository.CreateCompany(ctx, &company)
	if err != nil {
		return err
	}
	return nil
}

func (service *AgentService) FindCompanyByName(ctx context.Context, name string) *model.Company {
	company, _ := service.Repository.GetCompanyByName(ctx, name)
	return company
}

func (service *AgentService) ApproveCompany(ctx context.Context, dto *dto.ApproveCompanyDTO) {
	service.Repository.ApproveCompany(ctx, dto)
}

func (service *AgentService) GetAll(ctx context.Context) []*model.Company {
	return service.Repository.GetAll(ctx)
}
