package handlers

import (
	"XWS-Project/job_offer/service"
	pb "XWS-Project/proto/job_offer_service"
)

type JobHandler struct {
	pb.UnimplementedJobOfferServiceServer
	JobService *service.JobService
}
