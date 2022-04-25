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
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
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
	if err != nil {
		utilities.Tracer.LogError(span, err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	token, errr := utilities.CreateToken(user.ProfileID, "authentication_service")
	if errr != nil {
		utilities.Tracer.LogError(span, errr)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	responseDTO := dto.LoginResponseDTO{
		Token:     token,
		Email:     user.Email,
		Username:  user.Username,
		ProfileID: user.ProfileID,
	}

	respJson, err := json.Marshal(responseDTO)
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

func (handler *AuthenticationHandler) Preflight(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	rw.WriteHeader(http.StatusOK)
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
