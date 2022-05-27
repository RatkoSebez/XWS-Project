package handlers

import (
	"XWS-Project/job_offer/mapper"
	"XWS-Project/job_offer/service"
	pb "XWS-Project/proto/job_offer_service"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobHandler struct {
	pb.UnimplementedJobOfferServiceServer
	JobService *service.JobService
}

func (handler *JobHandler) CreateOffer(ctx context.Context, request *pb.OfferRequest) (*pb.Offer, error) {
	offer := mapper.MapCreateRequest(request)
	resp, err := handler.JobService.CreateOffer(ctx, *offer)
	response := mapper.MapOfferResponse(resp)
	return response, err
}

func (handler *JobHandler) GetOffersByCompany(ctx context.Context, request *pb.CompanyID) (*pb.Offers, error) {
	id, _ := primitive.ObjectIDFromHex(request.CompanyID)
	resp, err := handler.JobService.FindByCompany(ctx, id)
	response := mapper.MapOffersResponse(resp)
	return response, err
}

func (handler *JobHandler) GetOffersByPosition(ctx context.Context, request *pb.Position) (*pb.Offers, error) {
	resp, err := handler.JobService.FindByPosition(ctx, request.Position)
	response := mapper.MapOffersResponse(resp)
	return response, err
}
