package api

import (
	"XWS-Project/api_gateway/infrastructure/services"
	"XWS-Project/proto/login_service"
	//"XWS-Project/proto/post_service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type CustomHandler struct {
	postClientAddress           string
	authenticationClientAddress string
}

func NewCustomHandler(postClientAddress, authenticationClientAddress string) Handler {
	return &CustomHandler{
		postClientAddress:           postClientAddress,
		authenticationClientAddress: authenticationClientAddress,
	}
}

func (handler *CustomHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/login", handler.Login)
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
