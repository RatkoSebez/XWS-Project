package mapper

import "XWS-Project/follow/model"
import pb "XWS-Project/proto/follow_service"

func MapFollow(followDto *model.Follow) *pb.FollowMessage {
	rtrn := &pb.FollowMessage{
		Email: followDto.Email,
	}

	for _, el := range followDto.Followers {
		rtrn.Followers = append(rtrn.Followers, el)
	}
	for _, el := range followDto.Follows {
		rtrn.Follows = append(rtrn.Follows, el)
	}
	return rtrn
}

func MapFollowRequestMessage(req *pb.FollowRequestMessage) *model.FollowRequest {
	rtrn := &model.FollowRequest{
		ReceiverEmail: req.ReceiverEmail,
		SenderEmail:   req.SenderEmail,
	}
	return rtrn
}

func MapFollowRequest(req *model.FollowRequest) *pb.FollowRequestMessage {
	rtrn := &pb.FollowRequestMessage{
		ReceiverEmail: req.ReceiverEmail,
		SenderEmail:   req.SenderEmail,
	}
	return rtrn
}

func MapFollows(reqs []*model.FollowRequest) *pb.FollowRequestsMessage {
	rtrn := &pb.FollowRequestsMessage{
		Requests: nil,
	}

	for _, el := range reqs {
		rtrn.Requests = append(rtrn.Requests, MapFollowRequest(el))
	}
	return rtrn
}
