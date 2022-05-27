package mapper

import (
	"XWS-Project/job_offer/model"
	pb "XWS-Project/proto/job_offer_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCreateRequest(dto *pb.OfferRequest) *model.JobOffer {
	id, _ := primitive.ObjectIDFromHex(dto.CompanyId)
	resp := &model.JobOffer{
		CompanyId:      id,
		JobDescription: dto.JobDescription,
		JobOfferId:     primitive.NewObjectID(),
		Position:       dto.Position,
	}

	for _, el := range dto.Preconditions {
		resp.Preconditions = append(resp.Preconditions, el)
	}
	return resp
}

func MapOfferResponse(dto *model.JobOffer) *pb.Offer {
	resp := &pb.Offer{
		CompanyId:      dto.CompanyId.String(),
		JobOfferId:     dto.JobOfferId.String(),
		JobDescription: dto.JobDescription,
		Position:       dto.Position,
	}
	for _, el := range dto.Preconditions {
		resp.Preconditions = append(resp.Preconditions, el)
	}
	return resp
}

func MapOffersResponse(dto []*model.JobOffer) *pb.Offers {
	resp := &pb.Offers{
		Offers: nil,
	}
	for _, el := range dto {
		resp.Offers = append(resp.Offers, MapOfferResponse(el))
	}
	return resp

}
