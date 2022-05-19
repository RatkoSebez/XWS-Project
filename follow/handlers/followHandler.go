package handlers

import (
	"XWS-Project/follow/mapper"
	//"XWS-Project/follow/model"
	"XWS-Project/follow/service"
	pb "XWS-Project/proto/follow_service"
	//"XWS-Project/utilities"
	"context"
	//"encoding/json"
	//"github.com/gorilla/mux"
	//"net/http"
)

type FollowHandler struct {
	pb.UnimplementedFollowServiceServer
	FollowService *service.FollowService
}

/*func (handler *FollowHandler) GetFollow(rw http.ResponseWriter, r *http.Request) {
	span := utilities.Tracer.StartSpanFromRequest("follow-handler", r)
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)

	email, _ := mux.Vars(r)["email"]
	//fmt.Println("parameter: " + email)

	userFollows, _ := handler.FollowService.GetFollow(ctx, email)
	respJson, err := json.Marshal(userFollows)
	rw.Header().Add("content-type", "application/json")
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.Write(respJson)
}

func (handler *FollowHandler) CreateFollowRequest(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	span := utilities.Tracer.StartSpanFromRequest("follow-handler", r)
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)
	var followRequest model.FollowRequest
	err := json.NewDecoder(r.Body).Decode(&followRequest)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.FollowService.CreateFollowRequest(ctx, followRequest)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (handler *FollowHandler) GetFollowRequest(rw http.ResponseWriter, r *http.Request) {
	span := utilities.Tracer.StartSpanFromRequest("follow-handler", r)
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)

	email, _ := mux.Vars(r)["email"]
	//fmt.Println("parameter: " + email)

	followRequests, _ := handler.FollowService.GetFollowRequest(ctx, email)
	respJson, err := json.Marshal(followRequests)
	rw.Header().Add("content-type", "application/json")
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	rw.Write(respJson)
}

func (handler *FollowHandler) CreateFollow(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	span := utilities.Tracer.StartSpanFromRequest("follow-handler", r)
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)
	var followRequest model.FollowRequest
	err := json.NewDecoder(r.Body).Decode(&followRequest)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.FollowService.CreateFollow(ctx, followRequest)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}*/

//grpc handlers

func (handler *FollowHandler) GetFollow(ctx context.Context, pbDto *pb.EmptyEmailMessage) (*pb.FollowMessage, error) {
	userFollow, err := handler.FollowService.GetFollow(ctx, pbDto.Email)
	response := mapper.MapFollow(userFollow)
	return response, err
}

func (handler *FollowHandler) CreateFollow(ctx context.Context, pbDto *pb.FollowRequestMessage) (*pb.EmptyMessage, error) {
	request := mapper.MapFollowRequestMessage(pbDto)
	err := handler.FollowService.CreateFollow(ctx, *request)
	var response = &pb.EmptyMessage{}
	return response, err
}

func (handler *FollowHandler) GetFollowRequest(ctx context.Context, pbDto *pb.EmptyEmailMessage) (*pb.FollowRequestsMessage, error) {
	userFollows, _ := handler.FollowService.GetFollowRequest(ctx, pbDto.Email)
	response := mapper.MapFollows(userFollows)
	return response, nil
}

func (handler *FollowHandler) CreateFollowRequest(ctx context.Context, pbDto *pb.FollowRequestMessage) (*pb.EmptyMessage, error) {
	request := mapper.MapFollowRequestMessage(pbDto)
	err := handler.FollowService.CreateFollowRequest(ctx, *request)
	var response = &pb.EmptyMessage{}
	return response, err
}
