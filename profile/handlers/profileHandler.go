package handlers

import (
	"XWS-Project/profile/dto"
	"XWS-Project/profile/service"
	"XWS-Project/utilities"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type ProfileHandler struct {
	ProfileService *service.ProfileService
}

func (handler *ProfileHandler) Edit(w http.ResponseWriter, r *http.Request) {
	var infoDTO dto.ProfileDTO
	span := utilities.Tracer.StartSpanFromRequest("Profile-handler", r)
	err := json.NewDecoder(r.Body).Decode(&infoDTO)
	if err != nil {
		fmt.Println("greska1")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)
	vars := mux.Vars(r)
	var tempUsername = template.HTMLEscapeString(vars["username"])
	var user = handler.ProfileService.FindUserByUsername(ctx, infoDTO.Username, tempUsername)
	if user != nil && user.Username != tempUsername {
		fmt.Println("Username already taken")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.ProfileService.UpdateProfile(ctx, infoDTO, tempUsername)
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
