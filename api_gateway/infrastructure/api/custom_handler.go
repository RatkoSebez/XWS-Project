package api

import (
	"XWS-Project/api_gateway/infrastructure/services"
	"XWS-Project/proto/follow_service"
	"XWS-Project/proto/login_service"
	"XWS-Project/proto/post_service"
	"XWS-Project/proto/profile_service"
	"XWS-Project/proto/registration_service"
	"XWS-Project/utilities"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type CustomHandler struct {
	postClientAddress           string
	authenticationClientAddress string
	registrationClientAddress   string
	followClientAddress         string
	profileClientAddress        string
}

func NewCustomHandler(postClientAddress, authenticationClientAddress, regClientAdd, followClAdd, profileClAdd string) Handler {
	return &CustomHandler{
		postClientAddress:           postClientAddress,
		authenticationClientAddress: authenticationClientAddress,
		registrationClientAddress:   regClientAdd,
		followClientAddress:         followClAdd,
		profileClientAddress:        profileClAdd,
	}
}

func (handler *CustomHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/login", handler.Login)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/login", handler.Preflight)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/make-new-post", handler.MakePost)
	if err != nil {
		panic(err)
	}
	//	router.HandleFunc("/register", handler.Preflight).Methods("OPTIONS")
	err = mux.HandlePath("OPTIONS", "/register", handler.Preflight)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/save-photo", handler.SavePhoto)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/save-photo", handler.Preflight)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/comment", handler.Comment)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/comment", handler.Preflight)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/make-reaction", handler.MakeReaction)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/make-reaction", handler.Preflight)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/following-posts", handler.FollowingPosts)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/following-posts", handler.Preflight)
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
	err = mux.HandlePath("POST", "/follow/{email}", handler.CreateFollow)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/follow/{email}", handler.Preflight)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/followRequest/{email}", handler.GetFollowRequest)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("POST", "/followRequest/{email}", handler.CreateFollowRequest)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/followRequest/{email}", handler.Preflight)
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
	//err = mux.HandlePath("GET", "/users", handler.GetAllProfiles)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/users", handler.Preflight)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/{email}", handler.Preflight)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("GET", "/get-user-posts", handler.GetUserPosts)
	if err != nil {
		panic(err)
	}
	err = mux.HandlePath("OPTIONS", "/get-user-posts", handler.Preflight)
	if err != nil {
		panic(err)
	}

}

func (handler *CustomHandler) Test(w http.ResponseWriter, r *http.Request) {

}

func (handler *CustomHandler) Login(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,  OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,  OPTIONS, GET, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,  OPTIONS, GET, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,  OPTIONS, GET, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
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
		fmt.Println("GRESKA 1")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		fmt.Println("GRESKA 2")

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (handler *CustomHandler) GetUserPosts(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,  OPTIONS, GET, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	postClient := services.NewPostClient(handler.postClientAddress)
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	fmt.Println(userEmail)

	if userEmail == "" {
		fmt.Println("Nije ucitan mejl za postove")

		w.WriteHeader(http.StatusForbidden)
		return
	}
	var request post_service.UserMailMessage

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,  OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	fmt.Println("Registering user")
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

}

func (handler *CustomHandler) GetFollow(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	followClient := services.NewFollowClient(handler.followClientAddress)
	email := pathParams["email"]
	if email == "" {
		fmt.Println("Mail invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var request follow_service.FollowRequestMessage
	request.SenderEmail = userEmail
	request.ReceiverEmail = email
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	followClient := services.NewFollowClient(handler.followClientAddress)
	email := pathParams["email"]
	if email == "" {
		fmt.Println("Mail invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userEmail := utilities.GetLoggedUserEmailFromToken(r)
	if userEmail == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	var request follow_service.FollowRequestMessage
	request.SenderEmail = userEmail
	request.ReceiverEmail = email
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

// func (handler *CustomHandler) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", " OPTIONS,  GET")
// 	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
// 	profileClient := services.NewProfileClient(handler.profileClientAddress)
// 	//userEmail := utilities.GetLoggedUserEmailFromToken(r)

// 	var request profile_service.EmptyMailMessage

// 	response, err := profileClient.GetAllUsers(context.TODO(), &request)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	responseJson, err := json.Marshal(response)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(responseJson)
// }

func (handler *CustomHandler) GetProfileByMail(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, OPTIONS,  GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	profileClient := services.NewProfileClient(handler.profileClientAddress)
	//userEmail := utilities.GetLoggedUserEmailFromToken(r)
	email := pathParams["email"]
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// if userEmail == "" || userEmail != email {
	// 	fmt.Println("Nije ocitan header")
	// 	w.WriteHeader(http.StatusForbidden)
	// 	return
	// }
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
func (handler *CustomHandler) Preflight(rw http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	rw.WriteHeader(http.StatusOK)
}
