package api

import (
	"XWS-Project/api_gateway/infrastructure/services"
	"XWS-Project/proto/login_service"
	"XWS-Project/proto/post_service"
	"XWS-Project/proto/registration_service"
	"XWS-Project/utilities"
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type CustomHandler struct {
	postClientAddress           string
	authenticationClientAddress string
	registrationClientAddress   string
	followClientAddress         string
}

func NewCustomHandler(postClientAddress, authenticationClientAddress, regClientAdd, followClAdd string) Handler {
	return &CustomHandler{
		postClientAddress:           postClientAddress,
		authenticationClientAddress: authenticationClientAddress,
		registrationClientAddress:   regClientAdd,
		followClientAddress:         followClAdd,
	}
}

func (handler *CustomHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/login", handler.Login)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/make-new-post", handler.MakePost)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/save-photo", handler.SavePhoto)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/comment", handler.Comment)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/make-reaction", handler.MakeReaction)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/following-posts", handler.FollowingPosts)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/get-user-posts", handler.GetUserPosts)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/register", handler.GetUserPosts)
	if err != nil {
		panic(err)
	}
}

func (handler *CustomHandler) Test(w http.ResponseWriter, r *http.Request) {

}

func (handler *CustomHandler) Login(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	authClient := services.NewAuthenticationClient(handler.authenticationClientAddress)
	var request login_service.LoginDTO
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	postRequest := login_service.PostRequest{Dto: &request}
	loginResponse, err := authClient.Post(context.TODO(), &postRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	response, err := json.Marshal(loginResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (handler *CustomHandler) MakePost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	postClient := services.NewPostClient(handler.postClientAddress)

	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var request post_service.MakePostPlusEmail
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request.UserEmail = userEmail
	response, err := postClient.CreatePost(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) SavePhoto(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	postClient := services.NewPostClient(handler.postClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var request post_service.PhotoMessage
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request.UserEmail = userEmail
	response, err := postClient.SavePhoto(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) Comment(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	postClient := services.NewPostClient(handler.postClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var request post_service.NewComment
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request.UserEmail = userEmail
	response, err := postClient.Comment(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) MakeReaction(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	postClient := services.NewPostClient(handler.postClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var request post_service.NewReaction
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := postClient.MakeReaction(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) FollowingPosts(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	postClient := services.NewPostClient(handler.postClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var request post_service.GetListOfFollowing
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := postClient.FollowingPosts(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) GetUserPosts(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	postClient := services.NewPostClient(handler.postClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var request post_service.UserMailMessage
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request.UserMail = userEmail
	response, err := postClient.GetUserPosts(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) Register(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	registerClient := services.NewRegistrationClient(handler.registrationClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var request registration_service.RegisterMessage

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := registerClient.Register(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
