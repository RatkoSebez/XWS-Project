package handlers

import (
	"XWS-Project/registration/dto"
	"XWS-Project/registration/service"
	"XWS-Project/utilities"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegistrationHandler struct {
	RegistrationService *service.RegistrationService
}

func (handler *RegistrationHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerDTO dto.UserDTO
	span := utilities.Tracer.StartSpanFromRequest("Register-handler", r)
	err := json.NewDecoder(r.Body).Decode(&registerDTO)
	if err != nil {
		fmt.Println("greska1")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)

	var user = handler.RegistrationService.FindUserByUsername(ctx, registerDTO.Username)
	if user != nil {
		fmt.Println("Username already taken")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.RegistrationService.RegisterUser(ctx, registerDTO)
	if err != nil {
		//util.Tracer.LogError(span, err)
		fmt.Println("greska2")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte("{\"success\":\"ok\"}"))
	w.Header().Set("Content-Type", "application/json")
}
