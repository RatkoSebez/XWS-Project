package handlers

import (
	"XWS-Project/follow/model"
	"XWS-Project/follow/service"
	"XWS-Project/utilities"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type FollowHandler struct {
	FollowService *service.FollowService
}

func (handler *FollowHandler) GetFollow(rw http.ResponseWriter, r *http.Request) {
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
