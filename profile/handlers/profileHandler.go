package handlers

import (
	//"XWS-Project/profile/dto"
	"XWS-Project/profile/mapper"
	"XWS-Project/profile/service"
	pb "XWS-Project/proto/profile_service"
	"context"
	"net/http"
	//"github.com/gorilla/mux"
)

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	ProfileService *service.ProfileService
}

/*func (handler *ProfileHandler) Edit(w http.ResponseWriter, r *http.Request) {
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
	var tempMail = template.HTMLEscapeString(vars["email"])
	var user = handler.ProfileService.FindUserByMail(ctx, tempMail)

	var user2 = handler.ProfileService.FindUserByUsername(ctx, infoDTO.Username, user.Username)
	if user2 != nil && user2.Username != user.Username {
		fmt.Println("Username already taken")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.ProfileService.UpdateProfile(ctx, infoDTO, user.Username)
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

func (handler *ProfileHandler) GetProfileByMail(w http.ResponseWriter, r *http.Request) {
	//var infoDTO dto.ProfileDTO
	span := utilities.Tracer.StartSpanFromRequest("Profile-handler", r)

	vars := mux.Vars(r)
	var tempMail = template.HTMLEscapeString(vars["email"])
	ctx := utilities.Tracer.ContextWithSpan(context.Background(), span)
	var user = handler.ProfileService.FindUserByMail(ctx, tempMail)
	if user == nil {
		fmt.Println("User does not exist!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}*/
func (handler *ProfileHandler) Preflight(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST")
	rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
	rw.WriteHeader(http.StatusOK)
}

//grpc handlers

func (handler *ProfileHandler) Edit(ctx context.Context, request *pb.ProfileDTO) (*pb.StatusMessage, error) {
	infoDTO := mapper.MapProfileDTO(request)

	var tempMail = request.EmailParameter
	var user = handler.ProfileService.FindUserByMail(ctx, tempMail)
	var statusMesage = &pb.StatusMessage{
		Status: "",
	}

	var user2 = handler.ProfileService.FindUserByUsername(ctx, infoDTO.Username, user.Username)
	if user2 != nil && user2.Username != user.Username {
		statusMesage.Status = "Username already taken"
		return statusMesage, nil
	}

	err := handler.ProfileService.UpdateProfile(ctx, *infoDTO, user.Username)
	if err != nil {
		statusMesage.Status = "greska2"
		return statusMesage, err
	}
	statusMesage.Status = "Successfully edited profile."
	return statusMesage, nil
}

func (handler *ProfileHandler) GetProfileByMail(ctx context.Context, request *pb.EmptyMailMessage) (*pb.ProfileInfo, error) {
	var user = handler.ProfileService.FindUserByMail(ctx, request.Email)
	response := mapper.MapResponseProfile(user)
	return response, nil
}
