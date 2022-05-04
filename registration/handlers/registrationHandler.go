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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	var registerDTO dto.UserDTO
	span := utilities.Tracer.StartSpanFromRequest("Register-handler", r)
	err := json.NewDecoder(r.Body).Decode(&registerDTO)
	if err != nil {
		fmt.Println("GRESKA: najvrv los format jsona")
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
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("{\"success\":\"ok\"}"))
	w.Header().Set("Content-Type", "application/json")
}
func (handler *RegistrationHandler) Preflight(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	rw.WriteHeader(http.StatusOK)
}
