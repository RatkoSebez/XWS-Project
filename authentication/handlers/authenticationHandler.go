package handlers

import (
	"XWS-Project/authentication/dto"
	"XWS-Project/authentication/model"
	"XWS-Project/authentication/service"
	"XWS-Project/utilities"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthenticationHandler struct {
	AuthenticationService *service.AuthenticationService
}

func (handler *AuthenticationHandler) Login(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Loging in")
	var request dto.LoginDTO
	span := utilities.Tracer.StartSpanFromRequest("Login-handler", r)

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	var user *model.User
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)
	user, err = handler.AuthenticationService.Login(ctx, request)
	fmt.Print(user.Email)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (handler *AuthenticationHandler) Test(rw http.ResponseWriter, r *http.Request) {
	respJson, err := json.Marshal("vrati bilo sta")
	rw.Header().Add("content-type", "application/json")
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	_, _ = rw.Write(respJson)
}
