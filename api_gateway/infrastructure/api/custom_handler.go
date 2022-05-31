package api

import (
	"XWS-Project/api_gateway/infrastructure/services"
	"XWS-Project/proto/follow_service"
	"XWS-Project/proto/job_offer_service"
	"XWS-Project/proto/login_service"
	"XWS-Project/proto/post_service"
	"XWS-Project/proto/profile_service"
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
	profileClientAddress        string
	jobOfferClientAddress       string
}

func NewCustomHandler(postClientAddress, authenticationClientAddress, regClientAdd, followClAdd, profileClAdd, jobOfferClAdd string) Handler {
	return &CustomHandler{
		postClientAddress:           postClientAddress,
		authenticationClientAddress: authenticationClientAddress,
		registrationClientAddress:   regClientAdd,
		followClientAddress:         followClAdd,
		profileClientAddress:        profileClAdd,
		jobOfferClientAddress:       jobOfferClAdd,
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
	err = mux.HandlePath("POST", "/get-user-posts", handler.GetUserPosts)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/register", handler.Register)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/follow/{email}", handler.GetFollow)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/follow", handler.CreateFollow)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/followRequest/{email}", handler.GetFollowRequest)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/followRequest", handler.CreateFollowRequest)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("PUT", "/{email}", handler.Edit)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/{email}", handler.GetProfileByMail)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/create-offer", handler.CreateOffer)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/create-agent-offer", handler.CreateAgentOffer)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/get-by-coid", handler.GetOfferByCompanyID)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/get-by-position", handler.GetOfferByPosition)
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
	var request registration_service.RegisterMessage

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = registerClient.Register(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *CustomHandler) GetFollow(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	followClient := services.NewFollowClient(handler.followClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	email := pathParams["email"]
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var request follow_service.EmptyEmailMessage
	request.Email = email
	follow, err := followClient.GetFollow(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(follow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)

}

func (handler *CustomHandler) CreateFollow(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	followClient := services.NewFollowClient(handler.followClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var request follow_service.FollowRequestMessage
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	follow, err := followClient.CreateFollow(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(follow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) GetFollowRequest(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	followClient := services.NewFollowClient(handler.followClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	email := pathParams["email"]
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var request follow_service.EmptyEmailMessage
	request.Email = email
	follow, err := followClient.GetFollowRequest(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(follow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) CreateFollowRequest(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	followClient := services.NewFollowClient(handler.followClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var request follow_service.FollowRequestMessage
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	follow, err := followClient.CreateFollowRequest(context.TODO(), &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(follow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) Edit(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	profileClient := services.NewProfileClient(handler.profileClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	email := pathParams["email"]
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userEmail == "" || userEmail != email {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var request profile_service.ProfileDTO
	request.EmailParameter = email

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := profileClient.Edit(context.TODO(), &request)
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

func (handler *CustomHandler) GetProfileByMail(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	profileClient := services.NewProfileClient(handler.profileClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	email := pathParams["email"]
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if userEmail == "" || userEmail != email {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var request profile_service.EmptyMailMessage
	request.Email = email
	response, err := profileClient.GetProfileByMail(context.TODO(), &request)
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

func (handler *CustomHandler) CreateOffer(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	jobOfferClient := services.NewJobOfferClient(handler.jobOfferClientAddress)

	userEmail := utilities.GetLoggedUserEmailFromToken(r)

	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var request job_offer_service.OfferRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := jobOfferClient.CreateOffer(context.TODO(), &request)
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

func (handler *CustomHandler) GetOfferByCompanyID(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	jobOfferClient := services.NewJobOfferClient(handler.jobOfferClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var request job_offer_service.CompanyID
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := jobOfferClient.GetOffersByCompany(context.TODO(), &request)
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

func (handler *CustomHandler) GetOfferByPosition(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	jobOfferClient := services.NewJobOfferClient(handler.jobOfferClientAddress)

	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var request job_offer_service.Position
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := jobOfferClient.GetOffersByPosition(context.TODO(), &request)
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

func (handler *CustomHandler) CreateAgentOffer(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	jobOfferClient := services.NewJobOfferClient(handler.jobOfferClientAddress)
	var request job_offer_service.CreatOfferFromAgent
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	agentName := utilities.GetAgentNameFromPureToken(request.Token)
	if agentName == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if !handler.CheckDoesProfileExist(agentName) {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	response, err := jobOfferClient.CreateAgentOffer(context.TODO(), &request)
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

func (handler *CustomHandler) CheckDoesProfileExist(mail string) bool {
	profileClient := services.NewProfileClient(handler.profileClientAddress)
	var request profile_service.EmptyMailMessage
	request.Email = mail
	response, err := profileClient.GetProfileByMail(context.TODO(), &request)

	if err != nil {
		fmt.Println(err)
		return false
	}
	if response == nil {

		return false
	}
	return true
}
