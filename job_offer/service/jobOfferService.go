package service

import (
	"XWS-Project/job_offer/model"
	"XWS-Project/job_offer/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobService struct {
	JobRepo *repository.JobRepository
}

func (service *JobService) CreateOffer(ctx context.Context, offer model.JobOffer) (*model.JobOffer, error) {
	res, err := service.JobRepo.AddNew(ctx, offer)
	return res, err
}

func (service *JobService) FindByCompany(ctx context.Context, coId primitive.ObjectID) ([]*model.JobOffer, error) {
	res, err := service.JobRepo.FindByCompany(ctx, coId)
	return res, err
}

func (service *JobService) FindByPosition(ctx context.Context, position string) ([]*model.JobOffer, error) {
	res, err := service.JobRepo.FindByPosition(ctx, position)
	return res, err
}

func (service *JobService) FindAll(ctx context.Context) ([]*model.JobOffer, error) {
	res, err := service.JobRepo.GetAll(ctx)
	return res, err
}

func (service *JobService) Delete(ctx context.Context, id primitive.ObjectID) error {
	err := service.JobRepo.DeleteOffer(ctx, id)
	return err
}
