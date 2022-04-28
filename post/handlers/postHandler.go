package handlers

import (
	"XWS-Project/post/dto"
	"XWS-Project/post/model"
	"XWS-Project/post/service"
	"XWS-Project/utilities"
	"context"
	"encoding/json"
	"net/http"
)

type PostHandler struct {
	PostService *service.PostService
}

func (handler *PostHandler) MakePost(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	var request dto.NewPostDTO
	span := utilities.Tracer.StartSpanFromRequest("Post-handler", r)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var post *model.Post
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)
	userID := utilities.GetLoggedUserIDFromToken(r)
	post, err = handler.PostService.MakePost(ctx, request, userID)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		panic(err)
	}
	respJson, err := json.Marshal(post)
	if err != nil {
		utilities.Tracer.LogError(span, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(respJson)
	rw.Header().Set("Content-Type", "application/json")
	utilities.Tracer.FinishSpan(span)
}

func (handler *PostHandler) Preflight(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	rw.WriteHeader(http.StatusOK)
}

func (handler *PostHandler) SavePhoto(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	span := utilities.Tracer.StartSpanFromRequest("Photo-handler", r)
	var request dto.PhotoDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	userID := utilities.GetLoggedUserIDFromToken(r)
	handler.PostService.SavePhotos(request, userID)
	utilities.Tracer.FinishSpan(span)
}

func (handler *PostHandler) PostComment(rw http.ResponseWriter, r *http.Request) {
	var request dto.CommentDTO
	span := utilities.Tracer.StartSpanFromRequest("Post-handler", r)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	userID := utilities.GetLoggedUserIDFromToken(r)
	post, err := handler.PostService.PostComment(context.TODO(), request, userID)
	if err != nil {
		panic(err)
	}
	respJson, err := json.Marshal(post)
	if err != nil {
		utilities.Tracer.LogError(span, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(respJson)
	rw.Header().Set("Content-Type", "application/json")
	utilities.Tracer.FinishSpan(span)
}

func (handler *PostHandler) MakeReaction(rw http.ResponseWriter, r *http.Request) {

	var reaction dto.ReactionDTO
	span := utilities.Tracer.StartSpanFromRequest("React-to-the-post", r)

	err := json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	post, err := handler.PostService.MakeReaction(context.TODO(), reaction)
	if err != nil {
		panic(err)
	}
	respJson, err := json.Marshal(post)
	if err != nil {
		utilities.Tracer.LogError(span, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, _ = rw.Write(respJson)
	rw.Header().Set("Content-Type", "application/json")
	utilities.Tracer.FinishSpan(span)
}
